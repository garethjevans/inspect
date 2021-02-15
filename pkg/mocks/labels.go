package mocks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"testing"

	"github.com/sirupsen/logrus"
)

type MockLabelLister struct {
	Requests map[string]string
}

func (m *MockLabelLister) StubResponse(t *testing.T, repo string, tag string, file string) {
	if m.Requests == nil {
		m.Requests = make(map[string]string)
	}
	key := fmt.Sprintf("%s-%s", repo, tag)
	t.Logf("stubbing response for key '%s' with %s", key, file)
	m.Requests[key] = file
}

func (m *MockLabelLister) Labels(repo string, tag string) (map[string]string, error) {
	key := fmt.Sprintf("%s-%s", repo, tag)
	logrus.Infof("got request for key '%s'", key)

	file := m.Requests[key]

	data, err := ioutil.ReadFile(path.Join("testdata", file))
	if err != nil {
		return nil, err
	}

	labels := map[string]string{}
	err = json.Unmarshal(data, &labels)
	if err != nil {
		return nil, err
	}

	return labels, nil
}
