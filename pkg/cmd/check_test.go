package cmd_test

import (
	"testing"

	"github.com/garethjevans/inspect/pkg/mocks"

	"github.com/garethjevans/inspect/pkg/cmd"
	"github.com/stretchr/testify/assert"
)

func TestCheck(t *testing.T) {
	tests := []struct {
		image          string
		blobResponse   string
		expectedOutput string
	}{
		{
			image:        "jenkinscinfra/terraform",
			blobResponse: "blobs.1.0.0.json",
			expectedOutput: `+-----------------------------------+----+----------------+
| LABEL                             | OK | RECOMMENDATION |
+-----------------------------------+----+----------------+
| org.opencontainers.image.created  | OK |                |
| org.opencontainers.image.revision | OK |                |
| org.opencontainers.image.source   | OK |                |
| org.opencontainers.image.url      | OK |                |
| org.label-schema.build-date       | OK |                |
| org.label-schema.vcs-ref          | OK |                |
| org.label-schema.vcs-url          | OK |                |
| org.label-schema.url              | OK |                |
| inspect.tree.state                | OK |                |
+-----------------------------------+----+----------------+
`,
		},
		{
			image:        "jenkinscinfra/terraform",
			blobResponse: "blobs.no-labels.json",
			expectedOutput: `+-----------------------------------+---------+---------------------------------------------------------------------+
| LABEL                             | OK      | RECOMMENDATION                                                      |
+-----------------------------------+---------+---------------------------------------------------------------------+
| org.opencontainers.image.created  | Missing | date --utc +%Y-%m-%dT%H:%M:%S                                       |
| org.opencontainers.image.revision | Missing | git rev-parse --short HEAD                                          |
| org.opencontainers.image.source   | Missing | git config --get remote.origin.url                                  |
| org.opencontainers.image.url      | Missing | git config --get remote.origin.url                                  |
| org.label-schema.build-date       | Missing | date --utc +%Y-%m-%dT%H:%M:%S                                       |
| org.label-schema.vcs-ref          | Missing | git rev-parse --short HEAD                                          |
| org.label-schema.vcs-url          | Missing | git config --get remote.origin.url                                  |
| org.label-schema.url              | Missing | git config --get remote.origin.url                                  |
| inspect.tree.state                | Missing | test -z "$(git status --porcelain)" && echo "clean" || echo "dirty" |
+-----------------------------------+---------+---------------------------------------------------------------------+
`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.image, func(t *testing.T) {
			logger := &mocks.LoggerMock{}
			labelLister := mocks.MockLabelLister{}

			c := cmd.CheckCmd{
				BaseCmd: cmd.BaseCmd{
					Log: logger,
				},
				LabelLister: &labelLister,
			}

			cmd.Reset()

			labelLister.StubResponse(t, tc.image, "latest", tc.blobResponse)

			c.Args = []string{tc.image}

			err := c.Run()
			assert.NoError(t, err)

			assert.Equal(t, 1, len(logger.Messages))
			assert.Equal(t, tc.expectedOutput, logger.Messages[0])
		})
	}
}
