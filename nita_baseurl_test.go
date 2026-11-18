// Tests écrits à la main (NON générés) — protégés via .openapi-generator-ignore.
package nita

import "testing"

func TestBaseURLHTTPSObligatoire(t *testing.T) {
	acceptees := []string{
		"https://api.exemple.test",
		"HTTPS://api.exemple.test/base",
		"http://localhost:8190",
		"http://127.0.0.1:8190",
		"http://[::1]:8190",
	}
	for _, u := range acceptees {
		if _, err := resolveBaseURL(NitaConfig{BaseURL: u}); err != nil {
			t.Errorf("%q devrait être acceptée : %v", u, err)
		}
	}
	refusees := []string{
		"http://api.exemple.test",
		"http://localhost.exemple.test",
		"http://10.0.0.5:8190",
		"ftp://localhost",
		"api.exemple.test",
	}
	for _, u := range refusees {
		if _, err := resolveBaseURL(NitaConfig{BaseURL: u}); err == nil {
			t.Errorf("%q devrait être refusée", u)
		}
	}
}
