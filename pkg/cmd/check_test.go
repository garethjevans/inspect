package cmd_test

import (
	"testing"

	"github.com/garethjevans/inspect/pkg/cmd/mock"

	"github.com/garethjevans/inspect/pkg/inspect"
	"github.com/garethjevans/inspect/pkg/inspect/mocks"

	"github.com/garethjevans/inspect/pkg/cmd"
	"github.com/stretchr/testify/assert"
)

func TestCheck(t *testing.T) {
	tests := []struct {
		image string
		blobResponse         string
		expectedOutput string
	}{
		{
			image: "jenkinscinfra/terraform",
			blobResponse:        "blobs.1.0.0.json",
			expectedOutput: `+-----------------------------------+----+----------------+
| LABEL                             | OK | RECOMMENDATION |
+-----------------------------------+----+----------------+
| org.opencontainers.image.created  | OK |                |
| org.opencontainers.image.revision | OK |                |
| org.opencontainers.image.source   | OK |                |
| org.opencontainers.image.url      | OK |                |
+-----------------------------------+----+----------------+
`,
		},
		{
			image: "jenkinscinfra/terraform",
			blobResponse:        "blobs.no-labels.json",
			expectedOutput: `+-----------------------------------+---------+------------------------------------+
| LABEL                             | OK      | RECOMMENDATION                     |
+-----------------------------------+---------+------------------------------------+
| org.opencontainers.image.created  | Missing | date --utc +%Y-%m-%dT%H:%M:%S      |
| org.opencontainers.image.revision | Missing | git log -n 1 --pretty=format:%h    |
| org.opencontainers.image.source   | Missing | git config --get remote.origin.url |
| org.opencontainers.image.url      | Missing | git config --get remote.origin.url |
+-----------------------------------+---------+------------------------------------+
`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.image, func(t *testing.T) {
			logger := &mock.LoggerMock{}
			mock := &mocks.MockClient{}

			c := cmd.CheckCmd{
				BaseCmd: cmd.BaseCmd{
					Log: logger,
				},
				Client: inspect.Client{
					Client: mock,
				},
			}

			cmd.Reset()

			stubWithFixture(t, mock, "token.json")
			stubWithFixture(t, mock, "manifests.1.0.0.json")
			stubWithFixture(t, mock, tc.blobResponse)

			c.Args = []string{tc.image}

			err := c.Run()
			assert.NoError(t, err)

			assert.Equal(t, 1, len(logger.Messages))
			assert.Equal(t, tc.expectedOutput, logger.Messages[0])
		})
	}
}
