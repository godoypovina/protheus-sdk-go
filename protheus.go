package protheus

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	userAgent string = "Protheus Go SDK"
	apiURL    string = "http://192.168.1.247:8585"
	mimeJSON  string = "application/json"
)

// Protheus is the implementation to consume Protheus API services
type Protheus struct {
	Token string
}

// NewProtheus returns a new instance of the Protheus API services
func NewProtheus(token string) *Protheus {
	return &Protheus{
		Token: token,
	}
}

func (p *Protheus) get(resource string, params url.Values, dest interface{}) error {
	return p.doRequest("GET", resource, params, nil, dest)
}

func (p *Protheus) post(resource string, data interface{}, dest interface{}) error {
	buf, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return p.doRequest("POST", resource, nil, bytes.NewBuffer(buf), dest)
}

func (p *Protheus) put(resource string, data interface{}, dest interface{}) error {
	buf, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return p.doRequest("PUT", resource, nil, bytes.NewBuffer(buf), dest)
}

func (p *Protheus) newRequest(method string, uri string, body io.Reader) (*http.Request, error) {
	var buf []byte

	if body != nil {
		var err error
		buf, err = ioutil.ReadAll(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, uri, bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}

	if method == "POST" {
		req.Header.Set("Content-Type", mimeJSON)
		req.Header.Set("Accept", mimeJSON)
	}

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Authorization", "BASIC "+p.Token)

	return req, err
}

func (p *Protheus) doRequest(method string, resource string, params url.Values, body io.Reader, dest interface{}) error {
	//Build resource URL
	u, err := url.ParseRequestURI(apiURL)
	if err != nil {
		return err
	}
	u.Path = "/rest" + resource
	u.RawQuery = params.Encode()

	req, err := p.newRequest(method, u.String(), body)
	if err != nil {
		return err
	}

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(buf, dest); err != nil {
		return err
	}

	return nil
}
