// Couche mince écrite à la main (NON générée) — voir nita_client.go.
package nita

import (
	"encoding/base64"
	"encoding/json"
	"strings"
)

// jwtExp lit le champ exp (expiration, secondes epoch) d'un JWT SANS vérifier sa
// signature — suffisant pour décider s'il faut rafraîchir proactivement le token ;
// la vérification de signature est de toute façon faite côté serveur à chaque appel.
// ok=false si le token est malformé ou ne porte pas de exp.
func jwtExp(token string) (exp int64, ok bool) {
	parts := strings.Split(token, ".")
	if len(parts) < 2 {
		return 0, false
	}
	payload, err := decodeJWTSegment(parts[1])
	if err != nil {
		return 0, false
	}
	var claims struct {
		// float64 plutôt qu'int64 : robuste que le champ soit sérialisé comme entier
		// ou comme nombre à virgule.
		Exp float64 `json:"exp"`
	}
	if err := json.Unmarshal(payload, &claims); err != nil || claims.Exp == 0 {
		return 0, false
	}
	return int64(claims.Exp), true
}

// decodeJWTSegment décode un segment base64url de JWT, avec ou sans padding '='.
func decodeJWTSegment(seg string) ([]byte, error) {
	if m := len(seg) % 4; m != 0 {
		seg += strings.Repeat("=", 4-m)
	}
	return base64.URLEncoding.DecodeString(seg)
}
