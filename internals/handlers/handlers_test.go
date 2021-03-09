package handlers

import (
	"net/http"
	"net/url"
	"testing"
)


type postData struct {
	key   string
	value string
}

type routeTest struct {
	name               string
	url                string     // Route Url
	method             string     // HTTP method
	params             []postData // Any relevant post data
	expectedStatusCode int
}

var tests = []routeTest{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"ga", "/generals-quarter", "GET", []postData{}, http.StatusOK},
	{"mr", "/make-reservation", "GET", []postData{}, http.StatusOK},
	{"rs", "/reservation-summary", "GET", []postData{}, http.StatusOK},
	{"sa", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"post-sa", "/search-availability", "POST", []postData{
		{key: "start_date", value: "2020-01-01"},
		{key: "end_date", value: "2020-01-02"},
	}, http.StatusOK},
	{"post-mr", "/make-reservation", "POST", []postData{
		{key: "first_name", value: "Jon"},
		{key: "last_name", value: "Seow"},
		{key: "email", value: "test@test.com"},
		{key: "phone", value: "123-456-789"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	defer ts.Close()

	var resp *http.Response
	var err error

	for _, test := range tests {
		if test.method == "GET" {
			resp, err = ts.Client().Get(ts.URL + test.url)
		} else if test.method == "POST" {
			values := url.Values{}
			for _, data := range test.params {
				values.Add(data.key, data.value)
			}
			resp, err = ts.Client().PostForm(ts.URL + test.url, values)
		}

		if err != nil {
			t.Fatal(err)
		}

		if resp.StatusCode != test.expectedStatusCode {
			t.Errorf("Expected %d, but got %d", test.expectedStatusCode, resp.StatusCode)
		}
	}
}