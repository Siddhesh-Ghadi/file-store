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

func TestLsHandler(t *testing.T) {
    req1 := httptest.NewRequest(http.MethodGet, "http://localhost/ls", nil)
    w1 := httptest.NewRecorder()

    // test for success
    serverDir = "./testdata"
    lsHandler(w1, req1)
    if want1, got1 := http.StatusOK, w1.Result().StatusCode; want1 != got1 {
        t.Fatalf("expected: %v, instead got: %v", want1, got1)
    }

	want1 := `{"files":[{"name":"example.txt"},{"name":"sample.txt"}]}`
	resp1, _ := ioutil.ReadAll(w1.Result().Body)
	got1 := string(resp1)
	if want1 != got1 {
        t.Fatalf("expected: %v, instead got: %v", want1, got1)
    }

    // test for failure
    req2 := httptest.NewRequest(http.MethodGet, "http://localhost/ls", nil)
    w2 := httptest.NewRecorder()
    serverDir = "./testdat"
    lsHandler(w2, req2)
    if want2, got2 := http.StatusInternalServerError, w2.Result().StatusCode; want2 != got2 {
        t.Fatalf("expected: %v, instead got: %v", want2, got2)
    }

	want2 := `Something went wrong.`
	resp2, _ := ioutil.ReadAll(w2.Result().Body)
	got2 := string(resp2)
	if want2 != got2 {
        t.Fatalf("expected: %v, instead got: %v", want2, got2)
    }
}
