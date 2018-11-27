package stackongo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var host string = "http://api.stackexchange.com"
var transport http.RoundTripper

// UseSecure set to communicate using SSL
var UseSSL bool = true

type Session struct {
	Site string
}

func NewSession(site string) *Session {
	return &Session{Site: site}
}

func getTransport() http.RoundTripper {
	if transport != nil {
		return transport
	}
	return http.DefaultTransport
}

func SetTransport(t http.RoundTripper) {
	transport = t
}

// construct the endpoint URL
func setupEndpoint(path string, params map[string]string) *url.URL {
	base_url, _ := url.Parse(host)

	if UseSSL {
		base_url.Scheme = "https"
	} else {
		base_url.Scheme = "http"
	}

	endpoint, _ := base_url.Parse("/2.0/" + path)

	query := endpoint.Query()
	for key, value := range params {
		query.Set(key, value)
	}

	endpoint.RawQuery = query.Encode()

	return endpoint
}

// parse the response
func parseResponse(response *http.Response, result interface{}) error {
	// close the body when done reading
	defer response.Body.Close()

	//read the response
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, result)
	if err != nil {
		return err
	}

	if response.StatusCode == http.StatusBadRequest {
		return errors.New(fmt.Sprintf("Bad Request: %s", string(bytes)))
	}

	return nil
}

func (session Session) get(section string, params map[string]string, collection interface{}) (error error) {
	//set parameters for querystring
	params["site"] = session.Site

	return get(section, params, collection)
}

func get(section string, params map[string]string, collection interface{}) error {
	client := &http.Client{Transport: getTransport()}

	response, err := client.Get(setupEndpoint(section, params).String())
	if err != nil {
		return err
	}

	return parseResponse(response, collection)
}
