package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
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
	{"POST /search-availability", "/search-availability", "POST", []postData{
		{key: "start", value: "2020-01-01"},
		{key: "end", value: "2020-01-02"},
	}, http.StatusOK},
	{"POST /search-availability-json", "/search-availability-json", "POST", []postData{
		{key: "start", value: "2020-01-01"},
		{key: "end", value: "2020-01-02"},
	}, http.StatusOK},
	{"POST /make-reservation", "/make-reservation", "POST", []postData{
		{key: "first_name", value: "John"},
		{key: "last_name", value: "Smith"},
		{key: "email", value: "me@here.com"},
		{key: "phone", value: "555-555-5555"},
	}, http.StatusOK},
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
			// for POST requests
			values := url.Values{}
			for _, x := range e.params {
				values.Add(x.key, x.value)
			}
			resp, err := ts.Client().PostForm(ts.URL+e.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("Test %s: expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}