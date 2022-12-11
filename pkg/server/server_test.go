package server

import (
    "testing"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func TestRootHandler(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
    w := httptest.NewRecorder()
    rootHandler(w, req)
    // We should get a good status code
    if want, got := http.StatusOK, w.Result().StatusCode; want != got {
        t.Fatalf("expected: %v, instead got: %v", want, got)
    }

	want := "Welcome to File Store"
	resp, _ := ioutil.ReadAll(w.Result().Body)
	got := string(resp)
	if want != got {
        t.Fatalf("expected: %v, instead got: %v", want, got)
    }
}