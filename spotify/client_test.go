package spotify

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetTimeout(t *testing.T) {
	// Faux serveur lent : il dort 200 ms avant de répondre
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(200 * time.Millisecond)
	}))
	defer srv.Close()
	c := NewClient(50 * time.Millisecond)
	c.BaseURL = srv.URL
	var out map[string]any
	err := c.get("/n-importe-quoi", "tok", &out)
	if err == nil {
		t.Fatal("une erreur de timeout était attendue")
	}

}
