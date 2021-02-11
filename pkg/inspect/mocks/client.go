package mocks

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// MockClient is the mock client.
type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
	// Requests stores references to sequential requests that RoundTrip has received
	Requests      []*http.Request
	count         int
	responseStubs []*http.Response
}

var (
	// GetDoFunc fetches the mock client's `Do` func.
	GetDoFunc func(req *http.Request) (*http.Response, error)
)

// Do is the mock client's `Do` func.
func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	if len(m.responseStubs) <= m.count {
		return nil, fmt.Errorf("the MockClient: missing response stub for request %d, for url %s", m.count, req.URL.String())
	}
	resp := m.responseStubs[m.count]
	m.count++
	resp.Request = req
	m.Requests = append(m.Requests, req)
	return resp, nil
}

// StubResponse pre-records an HTTP response.
func (m *MockClient) StubResponse(status int, body io.Reader) {
	resp := &http.Response{
		StatusCode: status,
		Body:       ioutil.NopCloser(body),
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	m.responseStubs = append(m.responseStubs, resp)
}
