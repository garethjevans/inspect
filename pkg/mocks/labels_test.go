package mocks_test

import (
	"testing"

	"github.com/garethjevans/inspect/pkg/mocks"
	"github.com/stretchr/testify/assert"
)

func TestMockLabelLister(t *testing.T) {
	lister := mocks.MockLabelLister{}
	lister.StubResponse(t, "repo", "tag", "blobs.1.0.0.json")

	labels, err := lister.Labels("repo", "tag")
	assert.NoError(t, err)
	assert.Equal(t, 12, len(labels))
}
