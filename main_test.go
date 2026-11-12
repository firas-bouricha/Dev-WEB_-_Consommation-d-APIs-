package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthcheck(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	healthcheck(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("code = %d, attendu 200", rec.Code)
	}
	if rec.Body.String() != "ok" {
		t.Errorf("corps = %q, attendu \"ok\"", rec.Body.String())
	}
}
