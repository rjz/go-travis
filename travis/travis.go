package travis

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

const uaHeader = "go-travis/0.0.1"
const acceptHeader = "application/vnd.travis-ci.2+json"

type Client struct {
	apiHost *string
	token   *string
	client  *http.Client
}

func NewClient(accessToken *string) *Client {
	apiHost := "https://api.travis-ci.org"
	c := Client{&apiHost, accessToken, http.DefaultClient}
	return &c
}

func NewProClient(accessToken *string) *Client {
	apiHost := "https://api.travis-ci.com"
	c := Client{&apiHost, accessToken, http.DefaultClient}
	return &c
}

func (tc *Client) request(method, path string, body io.Reader) (*http.Request, error) {
	if tc.token == nil {
		panic("travis client is not configured.")
	}

	url := *tc.apiHost + path

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", acceptHeader)
	req.Header.Add("Authorization", "token "+*tc.token)
	req.Header.Add("User-agent", uaHeader)
	return req, nil
}

func (tc *Client) do(req *http.Request) ([]byte, error) {
	resp, err := tc.client.Do(req)
	if err != nil {
		return nil, err
	}
	data, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return nil, readErr
	}
	resp.Body.Close()
	if resp.StatusCode >= 300 { // We're not prepared to handle it
		return nil, errors.New(string(data))
	}
	return data, nil
}

func (tc *Client) Get(path string) ([]byte, error) {
	req, err := tc.request("GET", path, bytes.NewBuffer([]byte{}))
	if err != nil {
		return nil, err
	}
	return tc.do(req)
}

func (tc *Client) Post(path string, body interface{}) ([]byte, error) {
	jsonBytes, _ := json.Marshal(body)
	req, err := tc.request("POST", path, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	return tc.do(req)
}

func (tc *Client) Patch(path string, body interface{}) ([]byte, error) {
	jsonBytes, _ := json.Marshal(body)
	req, err := tc.request("PATCH", path, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	return tc.do(req)
}

func (tc *Client) Delete(path string) ([]byte, error) {
	req, err := tc.request("DELETE", path, bytes.NewBuffer([]byte{}))
	if err != nil {
		return nil, err
	}
	return tc.do(req)
}

func Bool(value bool) *bool {
	b := new(bool)
	*b = value
	return b
}

func String(value string) *string {
	b := new(string)
	*b = value
	return b
}
