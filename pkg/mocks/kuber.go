package mocks

import (
	"testing"
)

type MockImageLister struct {
	Requests map[string][]string
}

func (m *MockImageLister) StubResponse(t *testing.T, namespace string, values []string) {
	if m.Requests == nil {
		m.Requests = make(map[string][]string)
	}
	m.Requests[namespace] = values
}

func (m *MockImageLister) GetImagesForNamespace(namespace string) ([]string, error) {
	return m.Requests[namespace], nil
}
