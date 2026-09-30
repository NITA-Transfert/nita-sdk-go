// Tests écrits à la main (NON générés) : délai d'expiration HTTP du NitaClient.
package nita

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// serveurLent répond après délai (ou à la fermeture du serveur).
func serveurLent(t *testing.T, delai time.Duration, reponse string) *httptest.Server {
	t.Helper()
	fin := make(chan struct{})
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-time.After(delai):
		case <-fin:
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(reponse))
	}))
	t.Cleanup(func() {
		close(fin)
		srv.Close()
	})
	return srv
}

func TestConnectTimeout(t *testing.T) {
	srv := serveurLent(t, 5*time.Second, `{"data":{"token":"t"}}`)
	debut := time.Now()
	_, err := Connect(NitaConfig{BaseURL: srv.URL, APIKey: "k", HMACSecret: "s", Timeout: 200 * time.Millisecond})
	if err == nil {
		t.Fatal("attendu une erreur de délai")
	}
	if d := time.Since(debut); d > 2*time.Second {
		t.Fatalf("délai non respecté : %v", d)
	}
}

func TestAPIGenereeTimeout(t *testing.T) {
	// Authentification immédiate, puis API générée bloquée.
	fin := make(chan struct{})
	mux := http.NewServeMux()
	mux.HandleFunc("/api/authenticate", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"data":{"token":"t"}}`))
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-time.After(5 * time.Second):
		case <-fin:
		}
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()
	defer close(fin)

	nc, err := Connect(NitaConfig{BaseURL: srv.URL, APIKey: "k", Timeout: 200 * time.Millisecond})
	if err != nil {
		t.Fatal(err)
	}
	svc, ctx, err := nc.Compte()
	if err != nil {
		t.Fatal(err)
	}
	debut := time.Now()
	if _, _, err := svc.ConsulterSoldeCompte(ctx).Execute(); err == nil {
		t.Fatal("attendu une erreur de délai")
	}
	if d := time.Since(debut); d > 2*time.Second {
		t.Fatalf("délai non respecté : %v", d)
	}
}

func TestTimeoutParDefaut(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"data":{"token":"t"}}`))
	}))
	defer srv.Close()

	nc, err := Connect(NitaConfig{BaseURL: srv.URL, APIKey: "k"})
	if err != nil {
		t.Fatal(err)
	}
	if nc.httpClient.Timeout != DefaultTimeout {
		t.Fatalf("délai par défaut : %v", nc.httpClient.Timeout)
	}
	nc, err = Connect(NitaConfig{BaseURL: srv.URL, APIKey: "k", HTTPClient: &http.Client{Timeout: 7 * time.Second}})
	if err != nil {
		t.Fatal(err)
	}
	if nc.httpClient.Timeout != 7*time.Second {
		t.Fatalf("délai du HTTPClient fourni ignoré : %v", nc.httpClient.Timeout)
	}
}
