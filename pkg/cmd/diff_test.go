package cmd_test

import (
	"testing"

	"github.com/garethjevans/inspect/pkg/mocks"

	"github.com/garethjevans/inspect/pkg/cmd"
	"github.com/stretchr/testify/assert"
)

var (
	expectedDiffResponse = `+------------------------------------------+----------------------------------------------------------------+----------------------------------------------------------------+
| IMAGE                                    | 1                                                              | 2                                                              |
+------------------------------------------+----------------------------------------------------------------+----------------------------------------------------------------+
| jenkinsciinfra/terraform                 | 1.0.0                                                          | 1.1.0                                                          |
+------------------------------------------+----------------------------------------------------------------+----------------------------------------------------------------+
| inspect.tree.state                       | clean                                                          | clean                                                          |
| io.jenkins-infra.tools                   | golang,terraform                                               | golang,terraform                                               |
| io.jenkins-infra.tools.golang.version    | 1.15                                                           | 1.15                                                           |
| io.jenkins-infra.tools.terraform.version | 0.13.6                                                         | 0.13.6                                                         |
| org.label-schema.build-date              | 2021-01-27T08:34:20Z                                           | 2021-01-28T10:21:25Z                                           |
| org.label-schema.url                     | https://github.com/jenkins-infra/docker-terraform.git          | https://github.com/jenkins-infra/docker-terraform.git          |
| org.label-schema.vcs-ref                 | ad902ec                                                        | 441c261                                                        |
| org.label-schema.vcs-url                 | https://github.com/jenkins-infra/docker-terraform.git          | https://github.com/jenkins-infra/docker-terraform.git          |
| org.opencontainers.image.created         | 2021-01-27T08:34:20Z                                           | 2021-01-28T10:21:25Z                                           |
| org.opencontainers.image.revision        | ad902ec                                                        | 441c261                                                        |
| org.opencontainers.image.source          | https://github.com/jenkins-infra/docker-terraform.git          | https://github.com/jenkins-infra/docker-terraform.git          |
| org.opencontainers.image.url             | https://github.com/jenkins-infra/docker-terraform.git          | https://github.com/jenkins-infra/docker-terraform.git          |
+------------------------------------------+----------------------------------------------------------------+----------------------------------------------------------------+
| GitHub URL                               | https://github.com/jenkins-infra/docker-terraform/tree/ad902ec | https://github.com/jenkins-infra/docker-terraform/tree/441c261 |
+------------------------------------------+----------------------------------------------------------------+----------------------------------------------------------------+
`

	expectedDiffResponseMarkdown = `| Image | 1 | 2 |
| --- | --- | --- |
| jenkinsciinfra/terraform | 1.0.0 | 1.1.0 |
| inspect.tree.state | clean | clean |
| io.jenkins-infra.tools | golang,terraform | golang,terraform |
| io.jenkins-infra.tools.golang.version | 1.15 | 1.15 |
| io.jenkins-infra.tools.terraform.version | 0.13.6 | 0.13.6 |
| org.label-schema.build-date | 2021-01-27T08:34:20Z | 2021-01-28T10:21:25Z |
| org.label-schema.url | https://github.com/jenkins-infra/docker-terraform.git | https://github.com/jenkins-infra/docker-terraform.git |
| org.label-schema.vcs-ref | ad902ec | 441c261 |
| org.label-schema.vcs-url | https://github.com/jenkins-infra/docker-terraform.git | https://github.com/jenkins-infra/docker-terraform.git |
| org.opencontainers.image.created | 2021-01-27T08:34:20Z | 2021-01-28T10:21:25Z |
| org.opencontainers.image.revision | ad902ec | 441c261 |
| org.opencontainers.image.source | https://github.com/jenkins-infra/docker-terraform.git | https://github.com/jenkins-infra/docker-terraform.git |
| org.opencontainers.image.url | https://github.com/jenkins-infra/docker-terraform.git | https://github.com/jenkins-infra/docker-terraform.git |
| GitHub URL | https://github.com/jenkins-infra/docker-terraform/tree/ad902ec | https://github.com/jenkins-infra/docker-terraform/tree/441c261 |
`
)

func TestDiff(t *testing.T) {
	logger := &mocks.LoggerMock{}
	labelLister := &mocks.MockLabelLister{}

	c := cmd.DiffCmd{
		BaseCmd: cmd.BaseCmd{
			Log: logger,
		},
		LabelLister: labelLister,
	}

	cmd.Reset()

	labelLister.StubResponse(t, "jenkinsciinfra/terraform", "1.0.0", "blobs.1.0.0.json")
	labelLister.StubResponse(t, "jenkinsciinfra/terraform", "1.1.0", "blobs.1.1.0.json")

	c.Args = []string{"jenkinsciinfra/terraform:1.0.0", "jenkinsciinfra/terraform:1.1.0"}

	err := c.Run()
	assert.NoError(t, err)

	assert.Equal(t, 2, len(logger.Messages))
	assert.Equal(t, expectedDiffResponse, logger.Messages[0])
	assert.Equal(t, "https://github.com/jenkins-infra/docker-terraform/compare/ad902ec..441c261", logger.Messages[1])
}

func TestDiff_WithMarkdown(t *testing.T) {
	logger := &mocks.LoggerMock{}
	labelLister := &mocks.MockLabelLister{}

	c := cmd.DiffCmd{
		BaseCmd: cmd.BaseCmd{
			Log: logger,
		},
		LabelLister: labelLister,
	}

	cmd.Reset()
	cmd.EnableMarkdown()

	labelLister.StubResponse(t, "jenkinsciinfra/terraform", "1.0.0", "blobs.1.0.0.json")
	labelLister.StubResponse(t, "jenkinsciinfra/terraform", "1.1.0", "blobs.1.1.0.json")

	c.Args = []string{"jenkinsciinfra/terraform:1.0.0", "jenkinsciinfra/terraform:1.1.0"}

	err := c.Run()
	assert.NoError(t, err)

	assert.Equal(t, 2, len(logger.Messages))
	assert.Equal(t, expectedDiffResponseMarkdown, logger.Messages[0])
	assert.Equal(t, "https://github.com/jenkins-infra/docker-terraform/compare/ad902ec..441c261", logger.Messages[1])
}
