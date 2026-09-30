// Couche mince écrite à la main (NON générée) — protégée de la régénération via
// .openapi-generator-ignore (voir la liste en bas de ce fichier de commentaire dans
// .openapi-generator-ignore lui-même). Port Go de sdks/typescript/src/custom/ —
// même contrat (voir docs/superpowers/specs/2026-09-01-nita-client-couche-mince-
// multilangage-design.md dans NitaPlatform) : authentification (login -> JWT),
// rafraîchissement automatique du token, signature HMAC anti-rejeu et accesseurs
// par produit pré-configurés. Réutilise l'approche déjà validée dans
// smoke/go/main.go et examples/quickstart/go/internal/qsnita/nita.go
// (signingTransport + authenticate brut + contexte ContextAccessToken/ContextAPIKeys).
//
// Vit dans le MÊME package ("nita") que le SDK généré — pas de sous-paquet en Go,
// contrairement aux autres langages — donc pas d'import supplémentaire côté
// partenaire : `import nita "<module>/sdks/go"` suffit pour accéder à la fois aux
// types générés (APIClient, PartenaireToCashDtoV2…) et à NitaClient.
package nita

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

// --- environnements ---------------------------------------------------------

// URLs par environnement. Aucune n'est encore publique : tant qu'elles sont
// vides, le preset échoue avec un message explicite ; passer NitaConfig.BaseURL.
const (
	SandboxURL    = "" // TODO: URL publique du sandbox
	ProductionURL = "" // TODO: URL de production
)

// Presets acceptés par NitaConfig.Environment (voir NitaConfig.BaseURL, prioritaire).
const (
	EnvSandbox    = "sandbox"
	EnvProduction = "production"
)

func resolveBaseURL(cfg NitaConfig) (string, error) {
	if cfg.BaseURL != "" {
		return strings.TrimRight(cfg.BaseURL, "/"), nil
	}
	env := cfg.Environment
	if env == "" {
		env = EnvSandbox
	}
	var preset string
	switch env {
	case EnvSandbox:
		preset = SandboxURL
	case EnvProduction:
		preset = ProductionURL
	default:
		return "", fmt.Errorf("nita: environnement inconnu %q (%q|%q) ; passe NitaConfig.BaseURL explicitement", env, EnvSandbox, EnvProduction)
	}
	if preset == "" {
		return "", fmt.Errorf("nita: URL de base introuvable pour l'environnement %q ; passe NitaConfig.BaseURL explicitement (l'URL est indiquée dans ton espace partenaire)", env)
	}
	return preset, nil
}

// --- configuration ------------------------------------------------------------

// NitaConfig configure Connect.
type NitaConfig struct {
	// BaseURL explicite — prioritaire sur Environment.
	BaseURL string
	// Environment : preset (EnvSandbox par défaut, ou EnvProduction). Ignoré si
	// BaseURL est fourni.
	Environment string
	// APIKey partenaire (en-tête X-NT-API-KEY).
	APIKey string
	// Login / Password pour POST /api/authenticate.
	Login    string
	Password string
	// HMACSecret : si fourni, chaque requête à corps (hors multipart) est signée
	// (X-NT-TIMESTAMP / X-NT-NONCE / X-NT-SIGNATURE). Obligatoire en production ;
	// fixe en sandbox ("sandbox_test_signing_secret").
	HMACSecret string
	// HTTPClient : client de base optionnel (proxy, TLS, timeout, cookie jar…)
	// dont le Transport sera enveloppé par la signature HMAC. Si nil, un
	// *http.Client avec http.DefaultTransport est utilisé.
	HTTPClient *http.Client
}

// --- session (token + refreshToken) -------------------------------------------

type session struct {
	mu           sync.Mutex
	token        string
	refreshToken string
}

func (s *session) get() (token, refreshToken string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.token, s.refreshToken
}

func (s *session) set(token, refreshToken string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.token = token
	s.refreshToken = refreshToken
}

// refreshSkew : marge avant expiration à partir de laquelle le token est
// proactivement rafraîchi.
const refreshSkew = 30 * time.Second

// --- NitaClient ----------------------------------------------------------------

// NitaClient est la couche mince ergonomique au-dessus du SDK généré :
// authentification, rafraîchissement automatique du JWT, signature HMAC
// anti-rejeu (si HMACSecret est fourni) et accesseurs par produit
// pré-configurés (auth + signature déjà branchées).
//
// Usage :
//
//	nc, err := nita.Connect(nita.NitaConfig{
//		Environment: nita.EnvSandbox,
//		APIKey:      apiKey,
//		Login:       login,
//		Password:    password,
//		HMACSecret:  hmacSecret,
//	})
//	svc, ctx, err := nc.Transactions()
//	resp, _, err := svc.PartenaireToCash(ctx).PartenaireToCashDtoV2(dto).Execute()
//	data, err := nita.Unwrap(resp, resp.GetData(), err) // map[string]interface{}, ou *NitaError
//
// Chaque accesseur (Compte, Localites, Transactions, CheckStatus, MyNita,
// Achat) vérifie/rafraîchit le token AVANT de renvoyer le contexte — c'est donc
// l'appel à l'accesseur, pas Connect, qui déclenche le rafraîchissement proactif.
type NitaClient struct {
	baseURL    string
	apiKey     string
	hmacSecret string
	httpClient *http.Client // signingClient : Transport = &signingTransport{...}

	sess *session
	api  *APIClient // client généré, déjà configuré (Servers + HTTPClient)
}

// Connect authentifie contre l'API NITA (POST /api/authenticate — appel HTTP brut,
// cet endpoint est hors spec OpenAPI v2 donc absent du SDK généré), signé comme le
// reste si HMACSecret est fourni, puis construit le client généré (APIClient) câblé
// avec l'URL de base et le RoundTripper de signature.
func Connect(cfg NitaConfig) (*NitaClient, error) {
	baseURL, err := resolveBaseURL(cfg)
	if err != nil {
		return nil, err
	}
	if cfg.APIKey == "" {
		return nil, fmt.Errorf("nita: NitaConfig.APIKey manquant")
	}

	base := cfg.HTTPClient
	if base == nil {
		base = &http.Client{}
	}
	signingClient := &http.Client{
		Transport:     &signingTransport{secret: cfg.HMACSecret, base: base.Transport},
		CheckRedirect: base.CheckRedirect,
		Jar:           base.Jar,
		Timeout:       base.Timeout,
	}

	auth, err := postAuth(signingClient, baseURL, "/api/authenticate", cfg.APIKey, map[string]string{
		"username": cfg.Login,
		"password": cfg.Password,
	})
	if err != nil {
		return nil, err
	}

	apiCfg := NewConfiguration()
	apiCfg.Servers = ServerConfigurations{{URL: baseURL}}
	apiCfg.HTTPClient = signingClient

	return &NitaClient{
		baseURL:    baseURL,
		apiKey:     cfg.APIKey,
		hmacSecret: cfg.HMACSecret,
		httpClient: signingClient,
		sess:       &session{token: auth.token, refreshToken: auth.refreshToken},
		api:        NewAPIClient(apiCfg),
	}, nil
}

// Refresh force le rafraîchissement du token via POST /api/refreshToken. Appelé
// automatiquement par les accesseurs produit quand le JWT expire dans moins de 30s
// (voir validToken) ; exposé aussi publiquement pour un rafraîchissement manuel.
func (c *NitaClient) Refresh() error {
	_, refreshToken := c.sess.get()
	if refreshToken == "" {
		return fmt.Errorf("nita: aucun refreshToken disponible pour rafraîchir la session")
	}
	auth, err := postAuth(c.httpClient, c.baseURL, "/api/refreshToken", c.apiKey, map[string]string{
		"refreshToken": refreshToken,
	})
	if err != nil {
		return err
	}
	c.sess.set(auth.token, auth.refreshToken)
	return nil
}

// validToken renvoie un token valide, en rafraîchissant proactivement si le JWT
// expire dans moins de refreshSkew ET qu'un refreshToken est disponible — sinon
// renvoie le token courant tel quel (comme le TS de référence).
//
// Note : en cas d'appels concurrents proches de l'expiration, plusieurs
// rafraîchissements peuvent être déclenchés en parallèle (pas de verrou de
// coalescence) ; sans conséquence connue côté serveur mais documenté ici pour un
// futur durcissement si besoin.
func (c *NitaClient) validToken() (string, error) {
	token, refreshToken := c.sess.get()
	if exp, ok := jwtExp(token); ok && refreshToken != "" {
		if time.Until(time.Unix(exp, 0)) < refreshSkew {
			if err := c.Refresh(); err != nil {
				return "", err
			}
			token, _ = c.sess.get()
		}
	}
	return token, nil
}

// ctx construit le context.Context attendu par le SDK généré : le token (rafraîchi
// si besoin) sous ContextAccessToken, lu par APIClient.prepareRequest et ajouté en
// "Authorization: Bearer <token>" ; la clé API sous ContextAPIKeys (map, clé fixe
// "ApiKey"), lue par chaque Execute() et ajoutée en "X-NT-API-KEY".
func (c *NitaClient) ctx() (context.Context, error) {
	token, err := c.validToken()
	if err != nil {
		return nil, err
	}
	ctx := context.WithValue(context.Background(), ContextAccessToken, token)
	ctx = context.WithValue(ctx, ContextAPIKeys, map[string]APIKey{"ApiKey": {Key: c.apiKey}})
	return ctx, nil
}

// Ctx renvoie le context.Context prêt à l'emploi (token rafraîchi si besoin + clé
// API) — le même que celui renvoyé par chaque accesseur produit. Utile pour
// appeler directement un service généré non exposé par un accesseur produit.
func (c *NitaClient) Ctx() (context.Context, error) {
	return c.ctx()
}

// APIClient renvoie le client généré sous-jacent, déjà câblé (URL de base +
// signature HMAC) — échappatoire pour tout service non couvert par un accesseur
// produit. Combiner avec Ctx() pour l'authentification.
func (c *NitaClient) APIClient() *APIClient {
	return c.api
}

// --- accesseurs produit ---------------------------------------------------------
//
// Chacun renvoie le service généré correspondant PLUS un context.Context prêt à
// l'emploi (auth + clé API) ; l'obtention du contexte peut échouer (rafraîchissement
// réseau du token), d'où le 3e retour d'erreur — à vérifier avant d'appeler le
// service :
//
//	svc, ctx, err := nc.Compte()
//	if err != nil { ... }
//	resp, _, err := svc.ConsulterSoldeCompte(ctx).Execute()

// Compte renvoie CompteV2API (solde de compte) et son contexte.
func (c *NitaClient) Compte() (*CompteV2APIService, context.Context, error) {
	ctx, err := c.ctx()
	return c.api.CompteV2API, ctx, err
}

// Localites renvoie LocalitsV2API (villes, indicatifs) et son contexte.
func (c *NitaClient) Localites() (*LocalitsV2APIService, context.Context, error) {
	ctx, err := c.ctx()
	return c.api.LocalitsV2API, ctx, err
}

// Transactions renvoie TransactionsV2API (P2CASH, P2W, MyNita…) et
// son contexte.
func (c *NitaClient) Transactions() (*TransactionsV2APIService, context.Context, error) {
	ctx, err := c.ctx()
	return c.api.TransactionsV2API, ctx, err
}

// CheckStatus renvoie CheckingV2API (vérification de statut d'opération) et son
// contexte.
func (c *NitaClient) CheckStatus() (*CheckingV2APIService, context.Context, error) {
	ctx, err := c.ctx()
	return c.api.CheckingV2API, ctx, err
}

// MyNita renvoie MyNitaV2API (existence de compte) et son contexte.
func (c *NitaClient) MyNita() (*MyNitaV2APIService, context.Context, error) {
	ctx, err := c.ctx()
	return c.api.MyNitaV2API, ctx, err
}

// Achat renvoie AchatEnLigneV2API (achat en ligne) et son contexte.
func (c *NitaClient) Achat() (*AchatEnLigneV2APIService, context.Context, error) {
	ctx, err := c.ctx()
	return c.api.AchatEnLigneV2API, ctx, err
}

// P2p renvoie EnvoiInterPartenaireV2API (envoi inter-partenaire / P2P) et son contexte.
func (c *NitaClient) P2p() (*EnvoiInterPartenaireV2APIService, context.Context, error) {
	ctx, err := c.ctx()
	return c.api.EnvoiInterPartenaireV2API, ctx, err
}

// --- NewRequestID ----------------------------------------------------------------

// NewRequestID renvoie un requestId unique par appel : "<prefix>-<epoch>-<6 hex>".
// Un requestId réutilisé est REJETÉ côté serveur (pas de rejeu idempotent — P2C
// renvoie ERROR/400, P2W renvoie DUPLICATE/409) : chaque transaction a besoin d'un
// requestId neuf.
func NewRequestID(prefix string) string {
	return fmt.Sprintf("%s-%d-%s", prefix, time.Now().Unix(), randomHex(3))
}

// --- auth brute (hors spec v2) ---------------------------------------------------

// authSession est le résultat d'un POST /api/authenticate ou /api/refreshToken.
type authSession struct {
	token        string
	refreshToken string
}

// postAuth appelle path (POST, JSON, signé comme le reste si le client fourni
// signe — voir signingTransport) et lit {data:{token, refreshToken}} (avec repli
// sur {token} à la racine, forme alternative déjà tolérée côté TS). Ces deux
// endpoints sont hors spec OpenAPI v2, d'où l'appel HTTP brut plutôt qu'un service
// généré.
func postAuth(client *http.Client, baseURL, path, apiKey string, payload map[string]string) (authSession, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return authSession{}, err
	}
	req, err := http.NewRequest(http.MethodPost, baseURL+path, bytes.NewReader(body))
	if err != nil {
		return authSession{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-NT-API-KEY", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return authSession{}, err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return authSession{}, err
	}
	if resp.StatusCode >= 300 {
		return authSession{}, fmt.Errorf("nita: %s -> HTTP %d : %s", path, resp.StatusCode, truncate(string(respBody), 200))
	}

	var env struct {
		Token string `json:"token"`
		Data  struct {
			Token        string `json:"token"`
			RefreshToken string `json:"refreshToken"`
		} `json:"data"`
	}
	if err := json.Unmarshal(respBody, &env); err != nil {
		return authSession{}, fmt.Errorf("nita: réponse %s illisible : %w", path, err)
	}
	token := env.Data.Token
	if token == "" {
		token = env.Token
	}
	if token == "" {
		return authSession{}, fmt.Errorf("nita: JWT introuvable dans la réponse %s : %s", path, truncate(string(respBody), 200))
	}
	return authSession{token: token, refreshToken: env.Data.RefreshToken}, nil
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "…"
}
