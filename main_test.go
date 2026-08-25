package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	// Create a fake HTTP GET request for "/health".
	// NewRequest returns a *http.Request.
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	// Create a fake ResponseWriter.
	rec := httptest.NewRecorder()

	// Call the actual handler we want to test.
	// rec acts as the ResponseWriter (w)
	// req acts as the Request (r)
	healthHandler(rec, req)

	// rec.Code contains the status code written by healthHandler.
	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}
}

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	rec := httptest.NewRecorder()

	helloHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}
}
