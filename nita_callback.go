// Couche mince écrite à la main (NON générée) — protégée de la régénération via
// .openapi-generator-ignore. Voir nita_client.go pour le point d'entrée (Connect).
//
// Vérification d'un callback v2 signé envoyé par NITA au partenaire : POST JSON
// avec X-NT-TIMESTAMP, X-NT-NONCE et X-NT-SIGNATURE =
// hex(HMAC-SHA256(secret, ts + "\n" + nonce + "\n" + corps)). La signature porte
// sur les octets BRUTS reçus : ne jamais re-sérialiser le JSON avant vérification.
package nita

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// ErrInvalidCallback est enveloppée par toutes les erreurs de VerifyCallback
// (errors.Is(err, ErrInvalidCallback)).
var ErrInvalidCallback = errors.New("nita: callback invalide")

// DefaultCallbackTolerance est l'écart maximal admis entre l'horloge locale et
// X-NT-TIMESTAMP.
const DefaultCallbackTolerance = 300 * time.Second

// CallbackOptions configure VerifyCallback.
type CallbackOptions struct {
	// Secret HMAC du partenaire (obligatoire).
	Secret string
	// Tolerance : écart maximal d'horodatage (DefaultCallbackTolerance si <= 0).
	Tolerance time.Duration
	// IsNewNonce renvoie true si le nonce n'a jamais été vu, et le mémorise
	// pendant ttl (2 × Tolerance). Appelé uniquement si la signature est valide.
	// Si nil, AUCUN anti-rejeu : un callback valide rejoué dans la fenêtre de
	// tolérance est accepté. En production, brancher un stockage partagé entre
	// instances (base, Redis…).
	IsNewNonce func(nonce string, ttl time.Duration) bool
	// Now : horloge (time.Now si nil), injectable pour les tests.
	Now func() time.Time
}

func callbackErr(msg string) error {
	return fmt.Errorf("%w: %s", ErrInvalidCallback, msg)
}

// headerValue lit un en-tête sans tenir compte de la casse du nom : un
// http.Header construit à la main peut contenir des clés non canoniques.
func headerValue(header http.Header, name string) string {
	if v := header.Values(name); len(v) > 0 {
		return v[0]
	}
	for k, v := range header {
		if strings.EqualFold(k, name) && len(v) > 0 {
			return v[0]
		}
	}
	return ""
}

// VerifyCallback vérifie un callback NITA et renvoie son corps JSON (status,
// transaction_id…). rawBody doit être le corps tel que reçu (io.ReadAll(r.Body)).
func VerifyCallback(rawBody []byte, header http.Header, opts CallbackOptions) (map[string]any, error) {
	if opts.Secret == "" {
		return nil, callbackErr("Secret de signature manquant")
	}
	tolerance := opts.Tolerance
	if tolerance <= 0 {
		tolerance = DefaultCallbackTolerance
	}
	now := opts.Now
	if now == nil {
		now = time.Now
	}

	ts := headerValue(header, "X-NT-TIMESTAMP")
	nonce := headerValue(header, "X-NT-NONCE")
	signature := headerValue(header, "X-NT-SIGNATURE")
	if ts == "" || nonce == "" || signature == "" {
		return nil, callbackErr("En-tetes de signature manquants")
	}

	sec, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return nil, callbackErr("Horodatage hors tolerance")
	}
	ecart := now().Sub(time.Unix(sec, 0))
	if ecart < 0 {
		ecart = -ecart
	}
	if ecart > tolerance {
		return nil, callbackErr("Horodatage hors tolerance")
	}

	mac := hmac.New(sha256.New, []byte(opts.Secret))
	mac.Write([]byte(ts))
	mac.Write([]byte("\n"))
	mac.Write([]byte(nonce))
	mac.Write([]byte("\n"))
	mac.Write(rawBody)
	expected := hex.EncodeToString(mac.Sum(nil))
	if !hmac.Equal([]byte(expected), []byte(signature)) {
		return nil, callbackErr("Signature invalide")
	}

	if opts.IsNewNonce != nil && !opts.IsNewNonce(nonce, 2*tolerance) {
		return nil, callbackErr("Nonce deja utilise (rejeu)")
	}

	var payload map[string]any
	if err := json.Unmarshal(rawBody, &payload); err != nil || payload == nil {
		return nil, callbackErr("Corps JSON invalide")
	}
	return payload, nil
}
