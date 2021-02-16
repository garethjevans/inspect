package cmd_test

import (
	"testing"

	"github.com/garethjevans/inspect/pkg/mocks"

	"github.com/garethjevans/inspect/pkg/cmd"
	"github.com/stretchr/testify/assert"
)

func TestDiffNamespace(t *testing.T) {
	logger := &mocks.LoggerMock{}
	labelLister := &mocks.MockLabelLister{}
	imageLister := &mocks.MockImageLister{}

	c := cmd.DiffNamespaceCmd{
		BaseCmd: cmd.BaseCmd{
			Log: logger,
		},
		LabelLister: labelLister,
		ImageLister: imageLister,
	}

	cmd.Reset()

	imageLister.StubResponse(t, "namespace1", []string{"image1:1.0.0", "only-in-namespace1:1.0.0", "image2:1.0.0"})
	imageLister.StubResponse(t, "namespace2", []string{"only-in-namespace2:1.0.0", "image2:1.0.0", "image1:2.0.0"})

	c.Args = []string{"namespace1", "namespace2"}

	err := c.Run()
	assert.NoError(t, err)

	assert.Equal(t, 1, len(logger.Messages))
	assert.Equal(t, `+--------------------+------------+------------+
|                    | namespace1 | namespace2 |
+--------------------+------------+------------+
| image1             | 1.0.0      | 2.0.0      |
| image2             | 1.0.0      | 1.0.0      |
| only-in-namespace1 | 1.0.0      |            |
| only-in-namespace2 |            | 1.0.0      |
+--------------------+------------+------------+
`, logger.Messages[0])
}
