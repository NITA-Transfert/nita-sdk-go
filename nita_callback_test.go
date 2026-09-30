// Tests écrits à la main (NON générés) : vérification des callbacks v2 signés
// (vecteur commun aux 7 SDK).
package nita

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"net/http"
	"strings"
	"testing"
	"time"
)

const (
	cbSecret    = "s3cret"
	cbTimestamp = "1700000000"
	cbNonce     = "nonce-0123456789"
	cbBody      = `{"status":"1","transaction_id":"ACHAT1"}`
	cbSignature = "f4027a7eedeb961572515d3319ba9d78714689b687a96942f3eb0335abc61ec3"
)

func cbNow() time.Time { return time.Unix(1700000000, 0) }

func cbSign(ts, nonce, body string) string {
	mac := hmac.New(sha256.New, []byte(cbSecret))
	mac.Write([]byte(ts + "\n" + nonce + "\n" + body))
	return hex.EncodeToString(mac.Sum(nil))
}

func cbHeader(ts, nonce, sig string) http.Header {
	h := http.Header{}
	h.Set("X-NT-TIMESTAMP", ts)
	h.Set("X-NT-NONCE", nonce)
	h.Set("X-NT-SIGNATURE", sig)
	return h
}

func cbOpts() CallbackOptions {
	return CallbackOptions{Secret: cbSecret, Now: cbNow}
}

func expectRefus(t *testing.T, want string, body string, h http.Header, opts CallbackOptions) {
	t.Helper()
	_, err := VerifyCallback([]byte(body), h, opts)
	if err == nil {
		t.Fatalf("attendu %q, obtenu nil", want)
	}
	if !errors.Is(err, ErrInvalidCallback) {
		t.Fatalf("erreur non enveloppée dans ErrInvalidCallback : %v", err)
	}
	if !strings.HasSuffix(err.Error(), want) {
		t.Fatalf("attendu %q, obtenu %q", want, err.Error())
	}
}

func TestVerifyCallbackVecteurCommun(t *testing.T) {
	if got := cbSign(cbTimestamp, cbNonce, cbBody); got != cbSignature {
		t.Fatalf("vecteur commun : signature %s", got)
	}
	payload, err := VerifyCallback([]byte(cbBody), cbHeader(cbTimestamp, cbNonce, cbSignature), cbOpts())
	if err != nil {
		t.Fatal(err)
	}
	if payload["status"] != "1" || payload["transaction_id"] != "ACHAT1" {
		t.Fatalf("corps inattendu : %v", payload)
	}
}

func TestVerifyCallbackCasseQuelconque(t *testing.T) {
	// http.Header construit à la main : clés non canoniques.
	h := http.Header{
		"x-nt-timestamp": {cbTimestamp},
		"X-Nt-Nonce":     {cbNonce},
		"x-NT-signature": {cbSignature},
	}
	if _, err := VerifyCallback([]byte(cbBody), h, cbOpts()); err != nil {
		t.Fatal(err)
	}
}

func TestVerifyCallbackEnteteManquant(t *testing.T) {
	for _, name := range []string{"X-NT-TIMESTAMP", "X-NT-NONCE", "X-NT-SIGNATURE"} {
		h := cbHeader(cbTimestamp, cbNonce, cbSignature)
		h.Del(name)
		expectRefus(t, "En-tetes de signature manquants", cbBody, h, cbOpts())
	}
	expectRefus(t, "En-tetes de signature manquants", cbBody, cbHeader(cbTimestamp, "", cbSignature), cbOpts())
}

func TestVerifyCallbackSecretVide(t *testing.T) {
	opts := cbOpts()
	opts.Secret = ""
	expectRefus(t, "Secret de signature manquant", cbBody, cbHeader(cbTimestamp, cbNonce, cbSignature), opts)
}

func TestVerifyCallbackHorodatage(t *testing.T) {
	for _, ts := range []string{"1699999699", "1700000301"} {
		expectRefus(t, "Horodatage hors tolerance", cbBody, cbHeader(ts, cbNonce, cbSign(ts, cbNonce, cbBody)), cbOpts())
	}
	expectRefus(t, "Horodatage hors tolerance", cbBody, cbHeader("abc", cbNonce, cbSignature), cbOpts())
}

func TestVerifyCallbackSignatureModifiee(t *testing.T) {
	bad := cbSignature[:len(cbSignature)-1] + "0"
	expectRefus(t, "Signature invalide", cbBody, cbHeader(cbTimestamp, cbNonce, bad), cbOpts())
}

func TestVerifyCallbackCorpsModifie(t *testing.T) {
	body := strings.Replace(cbBody, `"1"`, `"2"`, 1)
	expectRefus(t, "Signature invalide", body, cbHeader(cbTimestamp, cbNonce, cbSignature), cbOpts())
}

func TestVerifyCallbackRejeu(t *testing.T) {
	vus := map[string]bool{}
	opts := cbOpts()
	opts.IsNewNonce = func(nonce string, ttl time.Duration) bool {
		if ttl != 600*time.Second {
			t.Fatalf("ttl attendu 600s, obtenu %v", ttl)
		}
		if vus[nonce] {
			return false
		}
		vus[nonce] = true
		return true
	}
	h := cbHeader(cbTimestamp, cbNonce, cbSignature)
	if _, err := VerifyCallback([]byte(cbBody), h, opts); err != nil {
		t.Fatal(err)
	}
	expectRefus(t, "Nonce deja utilise (rejeu)", cbBody, h, opts)
}

func TestVerifyCallbackNonceNonConsulteSiSignatureInvalide(t *testing.T) {
	appels := 0
	opts := cbOpts()
	opts.IsNewNonce = func(string, time.Duration) bool { appels++; return true }
	expectRefus(t, "Signature invalide", cbBody, cbHeader(cbTimestamp, cbNonce, strings.Repeat("0", 64)), opts)
	if appels != 0 {
		t.Fatalf("IsNewNonce appelé %d fois", appels)
	}
}

func TestVerifyCallbackCorpsNonJSON(t *testing.T) {
	for _, body := range []string{"pas du json", "[1,2]", "null"} {
		expectRefus(t, "Corps JSON invalide", body, cbHeader(cbTimestamp, cbNonce, cbSign(cbTimestamp, cbNonce, body)), cbOpts())
	}
}
