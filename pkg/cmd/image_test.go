package cmd_test

import (
	"testing"

	"github.com/garethjevans/inspect/pkg/mocks"

	"github.com/garethjevans/inspect/pkg/cmd"
	"github.com/stretchr/testify/assert"
)

var (
	expectedImageResponse = `+------------------------------------------+----------------------------------------------------------------+
| LABEL                                    | VALUE                                                          |
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
`

	expectedImageResponseRaw = ` LABEL                                     VALUE                                                          
 inspect.tree.state                        clean                                                          
 io.jenkins-infra.tools                    golang,terraform                                               
 io.jenkins-infra.tools.golang.version     1.15                                                           
 io.jenkins-infra.tools.terraform.version  0.13.6                                                         
 org.label-schema.build-date               2021-01-27T08:34:20Z                                           
 org.label-schema.url                      https://github.com/jenkins-infra/docker-terraform.git          
 org.label-schema.vcs-ref                  ad902ec                                                        
 org.label-schema.vcs-url                  https://github.com/jenkins-infra/docker-terraform.git          
 org.opencontainers.image.created          2021-01-27T08:34:20Z                                           
 org.opencontainers.image.revision         ad902ec                                                        
 org.opencontainers.image.source           https://github.com/jenkins-infra/docker-terraform.git          
 org.opencontainers.image.url              https://github.com/jenkins-infra/docker-terraform.git          
 GitHub URL                                https://github.com/jenkins-infra/docker-terraform/tree/ad902ec 
`

	expectedImageResponseMarkdown = `| Label | Value |
| --- | --- |
| inspect.tree.state | clean |
| io.jenkins-infra.tools | golang,terraform |
| io.jenkins-infra.tools.golang.version | 1.15 |
| io.jenkins-infra.tools.terraform.version | 0.13.6 |
| org.label-schema.build-date | 2021-01-27T08:34:20Z |
| org.label-schema.url | https://github.com/jenkins-infra/docker-terraform.git |
| org.label-schema.vcs-ref | ad902ec |
| org.label-schema.vcs-url | https://github.com/jenkins-infra/docker-terraform.git |
| org.opencontainers.image.created | 2021-01-27T08:34:20Z |
| org.opencontainers.image.revision | ad902ec |
| org.opencontainers.image.source | https://github.com/jenkins-infra/docker-terraform.git |
| org.opencontainers.image.url | https://github.com/jenkins-infra/docker-terraform.git |
| GitHub URL | https://github.com/jenkins-infra/docker-terraform/tree/ad902ec |
`

	expectedImageResponseNoHeaders = `+------------------------------------------+----------------------------------------------------------------+
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
`
)

func TestImage(t *testing.T) {
	logger := &mocks.LoggerMock{}
	labelLister := mocks.MockLabelLister{}

	c := cmd.ImageCmd{
		BaseCmd: cmd.BaseCmd{
			Log: logger,
		},
		LabelLister: &labelLister,
	}

	cmd.Reset()

	labelLister.StubResponse(t, "jenkinsciinfra/terraform", "1.0.0", "blobs.1.0.0.json")

	c.Args = []string{"jenkinsciinfra/terraform:1.0.0"}

	err := c.Run()
	assert.NoError(t, err)

	assert.Equal(t, 1, len(logger.Messages))
	assert.Equal(t, expectedImageResponse, logger.Messages[0])
}

func TestImage_Raw(t *testing.T) {
	logger := &mocks.LoggerMock{}
	labelLister := mocks.MockLabelLister{}

	c := cmd.ImageCmd{
		BaseCmd: cmd.BaseCmd{
			Log: logger,
		},
		LabelLister: &labelLister,
	}

	cmd.Reset()
	cmd.Raw()

	labelLister.StubResponse(t, "jenkinsciinfra/terraform", "1.0.0", "blobs.1.0.0.json")

	c.Args = []string{"jenkinsciinfra/terraform:1.0.0"}

	err := c.Run()
	assert.NoError(t, err)

	assert.Equal(t, 1, len(logger.Messages))
	assert.Equal(t, expectedImageResponseRaw, logger.Messages[0])
}

func TestImage_Markdown(t *testing.T) {
	logger := &mocks.LoggerMock{}
	labelLister := mocks.MockLabelLister{}

	c := cmd.ImageCmd{
		BaseCmd: cmd.BaseCmd{
			Log: logger,
		},
		LabelLister: &labelLister,
	}

	cmd.Reset()
	cmd.EnableMarkdown()

	labelLister.StubResponse(t, "jenkinsciinfra/terraform", "1.0.0", "blobs.1.0.0.json")

	c.Args = []string{"jenkinsciinfra/terraform:1.0.0"}

	err := c.Run()
	assert.NoError(t, err)

	assert.Equal(t, 1, len(logger.Messages))
	assert.Equal(t, expectedImageResponseMarkdown, logger.Messages[0])
}

func TestImage_NoHeaders(t *testing.T) {
	logger := &mocks.LoggerMock{}
	labelLister := mocks.MockLabelLister{}

	c := cmd.ImageCmd{
		BaseCmd: cmd.BaseCmd{
			Log: logger,
		},
		LabelLister: &labelLister,
	}

	cmd.Reset()
	cmd.DisableHeaders()

	labelLister.StubResponse(t, "jenkinsciinfra/terraform", "1.0.0", "blobs.1.0.0.json")

	c.Args = []string{"jenkinsciinfra/terraform:1.0.0"}

	err := c.Run()
	assert.NoError(t, err)

	assert.Equal(t, 1, len(logger.Messages))
	assert.Equal(t, expectedImageResponseNoHeaders, logger.Messages[0])
}

func TestImage_NoLabels(t *testing.T) {
	logger := &mocks.LoggerMock{}
	labelLister := mocks.MockLabelLister{}

	c := cmd.ImageCmd{
		BaseCmd: cmd.BaseCmd{
			Log: logger,
		},
		LabelLister: &labelLister,
	}

	cmd.Reset()

	labelLister.StubResponse(t, "jenkinsciinfra/terraform", "1.0.0", "blobs.no-labels.json")

	c.Args = []string{"jenkinsciinfra/terraform:1.0.0"}

	err := c.Run()
	assert.NoError(t, err)

	assert.Equal(t, 1, len(logger.Messages))
	assert.Equal(t, "No labels found for jenkinsciinfra/terraform:1.0.0", logger.Messages[0])
}
