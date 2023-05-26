package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	wantCode := http.StatusOK
	wantBody := "Hello, Nackademin!"

	srv := &server{
		httpServer: &http.Server{
			Addr: "localhost:8080",
		},
	}

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	srv.helloHandler().ServeHTTP(rr, req)

	if wantCode != rr.Code {
		t.Errorf("helloHandler() = unexpected result, want: %d, got: %d\n", wantCode, rr.Code)
	}

	if wantBody != rr.Body.String() {
		t.Errorf("helloHandler() = unexpected result, want: %s, got: %s\n", wantBody, rr.Body.String())
	}
}
