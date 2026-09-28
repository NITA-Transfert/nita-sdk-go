// Couche mince écrite à la main (NON générée) — protégée de la régénération via
// .openapi-generator-ignore. Voir nita_client.go pour le point d'entrée (Connect).
//
// Signature HMAC anti-rejeu, obligatoire en production, optionnelle en sandbox.
// Formule (identique prod/sandbox) : hex(HMAC-SHA256(secret, ts + "\n" + nonce + "\n" + corps)).
// Le séparateur est "\n" (0x0A), jamais "\r\n". Reprise telle quelle de
// smoke/go/main.go et examples/quickstart/go/internal/qsnita/nita.go.
package nita

import (
	"bytes"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// signHeaders calcule les 3 en-têtes X-NT-* pour le corps donné. Doit être appelé sur
// les octets réellement émis sur le fil (voir signingTransport).
func signHeaders(secret string, body []byte) map[string]string {
	ts := strconv.FormatInt(time.Now().Unix(), 10)
	nonce := randomHex(16)
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(ts))
	mac.Write([]byte("\n"))
	mac.Write([]byte(nonce))
	mac.Write([]byte("\n"))
	mac.Write(body)
	sig := hex.EncodeToString(mac.Sum(nil))
	return map[string]string{
		"X-NT-TIMESTAMP": ts,
		"X-NT-NONCE":     nonce,
		"X-NT-SIGNATURE": sig,
	}
}

// randomHex renvoie n octets aléatoires cryptographiques encodés en hexadécimal.
func randomHex(n int) string {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err) // crypto/rand qui échoue est irrécupérable
	}
	return hex.EncodeToString(b)
}

// isMultipart indique si un Content-Type est une requête multipart — jamais signée
// (le corps n'est pas un texte stable/rejouable de la même façon qu'un JSON).
func isMultipart(contentType string) bool {
	return strings.HasPrefix(strings.ToLower(strings.TrimSpace(contentType)), "multipart/")
}

// signingTransport est l'intercepteur HTTP (http.RoundTripper) qui signe chaque
// requête à corps non-multipart : il tamponne req.Body, signe EXACTEMENT ces octets,
// pose les 3 en-têtes X-NT-*, restaure le corps puis délègue à base
// (http.DefaultTransport si nil). Si secret est vide, ne signe rien — utile pour un
// NitaClient sans hmacSecret (sandbox sans anti-rejeu).
//
// Câblé via generatedConfiguration.HTTPClient = &http.Client{Transport: &signingTransport{...}}
// dans Connect (nita_client.go). Identique à l'approche validée dans smoke/go/main.go.
type signingTransport struct {
	secret string
	base   http.RoundTripper
}

func (t *signingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	base := t.base
	if base == nil {
		base = http.DefaultTransport
	}
	if t.secret == "" {
		return base.RoundTrip(req)
	}

	var bodyBytes []byte
	if req.Body != nil {
		var err error
		bodyBytes, err = io.ReadAll(req.Body)
		req.Body.Close()
		if err != nil {
			return nil, fmt.Errorf("nita: lecture du corps pour signature : %w", err)
		}
		req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		req.ContentLength = int64(len(bodyBytes))
	}

	// Le multipart n'est pas signé (le serveur ne vérifie pas sa signature).
	// TOUTES les autres requêtes — Y COMPRIS les GET sans corps — sont signées
	// (corps vide => signature sur []byte vide) : le sandbox exige les 3
	// en-têtes X-NT-* sur chaque requête.
	if isMultipart(req.Header.Get("Content-Type")) {
		return base.RoundTrip(req)
	}
	for k, v := range signHeaders(t.secret, bodyBytes) {
		req.Header.Set(k, v)
	}
	return base.RoundTrip(req)
}
