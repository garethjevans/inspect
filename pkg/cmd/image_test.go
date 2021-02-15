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
| jenkinsciinfra/terraform                 | 1.0.0                                                          |
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
 jenkinsciinfra/terraform                  1.0.0                                                          
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
| jenkinsciinfra/terraform | 1.0.0 |
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
| jenkinsciinfra/terraform                 | 1.0.0                                                          |
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
)

func TestImage(t *testing.T) {
	tests := []struct {
		name           string
		labelFile      string
		expectedOutput string
		testFunc       func()
	}{
		{
			name:           "default",
			labelFile:      "blobs.1.0.0.json",
			expectedOutput: expectedImageResponse,
		},
		{
			name:           "raw",
			labelFile:      "blobs.1.0.0.json",
			expectedOutput: expectedImageResponseRaw,
			testFunc: func() {
				cmd.Raw()
			},
		},
		{
			name:           "markdown",
			labelFile:      "blobs.1.0.0.json",
			expectedOutput: expectedImageResponseMarkdown,
			testFunc: func() {
				cmd.EnableMarkdown()
			},
		},
		{
			name:           "no-headers",
			labelFile:      "blobs.1.0.0.json",
			expectedOutput: expectedImageResponseNoHeaders,
			testFunc: func() {
				cmd.DisableHeaders()
			},
		},
		{
			name:           "no-labels",
			labelFile:      "blobs.no-labels.json",
			expectedOutput: "No labels found for jenkinsciinfra/terraform:1.0.0",
			testFunc: func() {
				cmd.DisableHeaders()
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			logger := &mocks.LoggerMock{}
			labelLister := mocks.MockLabelLister{}

			c := cmd.ImageCmd{
				BaseCmd: cmd.BaseCmd{
					Log: logger,
				},
				LabelLister: &labelLister,
			}

			cmd.Reset()

			if tc.testFunc != nil {
				tc.testFunc()
			}

			labelLister.StubResponse(t, "jenkinsciinfra/terraform", "1.0.0", tc.labelFile)

			c.Args = []string{"jenkinsciinfra/terraform:1.0.0"}

			err := c.Run()
			assert.NoError(t, err)

			assert.Equal(t, 1, len(logger.Messages))
			assert.Equal(t, tc.expectedOutput, logger.Messages[0])
		})
	}
}
