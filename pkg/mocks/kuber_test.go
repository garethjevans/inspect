package mocks_test

import (
	"testing"

	"github.com/garethjevans/inspect/pkg/mocks"
	"github.com/stretchr/testify/assert"
)

func TestMockImageLister(t *testing.T) {
	lister := mocks.MockImageLister{}
	lister.StubResponse(t, "namespace", []string{"image:one", "image:two", "image:3"})

	images, err := lister.GetImagesForNamespace("namespace")
	assert.NoError(t, err)
	assert.Equal(t, 3, len(images))
}
