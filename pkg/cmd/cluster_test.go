package cmd_test

import (
	"fmt"
	"testing"

	"github.com/garethjevans/inspect/pkg/mocks"

	"github.com/garethjevans/inspect/pkg/cmd"
	"github.com/stretchr/testify/assert"
)

func TestCluster(t *testing.T) {
	tests := []struct {
		namespace string
		images    []string
		output    []string
	}{
		{
			namespace: "",
			images:    []string{"jenkins/jenkins:lts"},
			output: []string{`+------------------------------------------+----------------------------------------------------------------+
| LABEL                                    | VALUE                                                          |
+------------------------------------------+----------------------------------------------------------------+
| jenkins/jenkins                          | lts                                                            |
+------------------------------------------+----------------------------------------------------------------+
| inspect.tree.state                       | clean                                                          |
| io.jenkins-infra.tools                   | golang,terraform                                               |
| io.jenkins-infra.tools.golang.version    | 1.15                                                           |
| io.jenkins-infra.tools.terraform.version | 0.13.6                                                         |
| org.label-schema.build-date              | 2021-01-27T08:34:20Z                                           |
| org.label-schema.url                     | https://github.com/jenkins-infra/docker-terraform.git          |
| org.label-schema.vcs-ref                 | ad902ec                                                        |
| org.label-schema.vcs-url                 | https://github.com/jenkins-infra/docker-terraform.git          |
| org.opencontainers.image.created         | 2021-01-27T08:34:20Z                                           |
| org.opencontainers.image.revision        | ad902ec                                                        |
| org.opencontainers.image.source          | https://github.com/jenkins-infra/docker-terraform.git          |
| org.opencontainers.image.url             | https://github.com/jenkins-infra/docker-terraform.git          |
+------------------------------------------+----------------------------------------------------------------+
| GitHub URL                               | https://github.com/jenkins-infra/docker-terraform/tree/ad902ec |
+------------------------------------------+----------------------------------------------------------------+
`},
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%s", tc.images), func(t *testing.T) {
			logger := &mocks.LoggerMock{}
			labelLister := mocks.MockLabelLister{}
			imageLister := mocks.MockImageLister{}

			c := cmd.ClusterCmd{
				BaseCmd: cmd.BaseCmd{
					Log: logger,
				},
				LabelLister: &labelLister,
				ImageLister: &imageLister,
			}

			c.Namespace = tc.namespace

			cmd.Reset()

			imageLister.StubResponse(t, tc.namespace, tc.images)
			labelLister.StubResponse(t, "jenkins/jenkins", "lts", "blobs.1.0.0.json")

			err := c.Run()
			assert.NoError(t, err)

			assert.Equal(t, len(tc.output), len(logger.Messages))
			assert.Equal(t, tc.output, logger.Messages)
		})
	}
}
