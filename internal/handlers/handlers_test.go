package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"GET home", "/", "GET", []postData{}, http.StatusOK},
	{"GET /about", "/about", "GET", []postData{}, http.StatusOK},
	{"GET /generals-quarters", "/generals-quarters", "GET", []postData{}, http.StatusOK},
	{"GET /majors-suite", "/majors-suite", "GET", []postData{}, http.StatusOK},
	{"GET /search-availability", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"GET /contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"GET /make-reservation", "/make-reservation", "GET", []postData{}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("Test %s: expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		} else {
		}
	}
}
