package spotify

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestFetchToken(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "mauvaise méthode", 405)
			return
		}
		r.ParseForm()
		if r.PostForm.Get("grant_type") != "client_credentials" {
			http.Error(w, "mauvais grant_type", 400)
			return
		}
		if r.PostForm.Get("client_id") != "id-test" || r.PostForm.Get("client_secret") != "secret-test" {
			http.Error(w, "identifiants manquants", 400)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"access_token":"faux-token","token_type":"Bearer","expires_in":3600}`)
	}))
	defer srv.Close()

	c := NewClient(2 * time.Second)
	c.TokenURL = srv.URL
	c.ClientID = "id-test"
	c.ClientSecret = "secret-test"
	tok, err := c.FetchToken()
	if err != nil {
		t.Fatalf("FetchToken: %v", err)
	}
	if tok.AccessToken != "faux-token" {
		t.Errorf("AccessToken: got %q, want %q", tok.AccessToken, "faux-token")
	}
	if tok.ExpiresIn != 3600 {
		t.Errorf("ExpiresIn: got %d, want %d", tok.ExpiresIn, 3600)
	}
}
