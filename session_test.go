package stackongo

import (
	"testing"
	"http"
	"http/httptest"
)

func returnDummyResponseForPath(path string, dummy_response string, t *testing.T) *httptest.Server {
	//serve dummy responses
	dummy_data := []byte(dummy_response)
	dummy_server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != path {
			t.Error("Invalid path")
		}
		w.Write(dummy_data)
	}))

	//change the host to use the test server
	setHost(dummy_server.URL)

	return dummy_server
}
