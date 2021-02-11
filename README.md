[![Go Report Card](https://goreportcard.com/badge/github.com/garethjevans/inspect)](https://goreportcard.com/report/github.com/garethjevans/inspect)
[![Downloads](https://img.shields.io/github/downloads/garethjevans/inspect/total.svg)]()

# inspect

a small CLI can query the metadata of an image to try to determine its origin.

currrently only works with images stored on dockerhub with their source stored on GitHub.

## To Install

```
brew tap garethjevans/tap
brew install inspect
```

This can be used a docker container with the following:

```
docker run -it garethjevans/inspect
```

## Usage

To inspect an image

```
inspect image <image>
```
e.g.


```
inspect image jenkinsciinfra/terraform
```

Will produce an output similar to:

```
+------------------------------------------+----------------------------------------------------------------+
| LABEL                                    | VALUE                                                          |
+------------------------------------------+----------------------------------------------------------------+
| org.opencontainers.image.created         | 2021-02-05T18:16:06Z                                           |
| org.opencontainers.image.revision        | d25f040                                                        |
| org.opencontainers.image.source          | https://github.com/jenkins-infra/docker-terraform.git          |
| io.jenkins-infra.tools.golang.version    | 1.15                                                           |
| org.label-schema.url                     | https://github.com/jenkins-infra/docker-terraform.git          |
| org.label-schema.vcs-url                 | https://github.com/jenkins-infra/docker-terraform.git          |
| org.label-schema.vcs-ref                 | d25f040                                                        |
| org.opencontainers.image.url             | https://github.com/jenkins-infra/docker-terraform.git          |
| io.jenkins-infra.tools                   | golang,terraform                                               |
| io.jenkins-infra.tools.terraform.version | 0.13.6                                                         |
| org.label-schema.build-date              | 2021-02-05T18:16:06Z                                           |
+------------------------------------------+----------------------------------------------------------------+
| GitHub URL                               | https://github.com/jenkins-infra/docker-terraform/tree/d25f040 |
+------------------------------------------+----------------------------------------------------------------+
```

To perform a diff between two images: 

```
inspect diff jenkinsciinfra/terraform:1.0.0 jenkinsciinfra/terraform:1.1.0
```

Will produce output like:

```
+------------------------------------------+----------------------------------------------------------------+----------------------------------------------------------------+
| IMAGE                                    | 1                                                              | 2                                                              |
+------------------------------------------+----------------------------------------------------------------+----------------------------------------------------------------+
| jenkinsciinfra/terraform                 | 1.0.0                                                          | 1.1.0                                                          |
+------------------------------------------+----------------------------------------------------------------+----------------------------------------------------------------+
| io.jenkins-infra.tools.terraform.version | 0.13.6                                                         | 0.13.6                                                         |
| org.label-schema.build-date              | 2021-01-27T08:34:20Z                                           | 2021-01-28T10:21:25Z                                           |
| org.label-schema.vcs-ref                 | ad902ec                                                        | 441c261                                                        |
| org.label-schema.vcs-url                 | https://github.com/jenkins-infra/docker-terraform.git          | https://github.com/jenkins-infra/docker-terraform.git          |
| org.opencontainers.image.created         | 2021-01-27T08:34:20Z                                           | 2021-01-28T10:21:25Z                                           |
| org.opencontainers.image.source          | https://github.com/jenkins-infra/docker-terraform.git          | https://github.com/jenkins-infra/docker-terraform.git          |
| io.jenkins-infra.tools                   | golang,terraform                                               | golang,terraform                                               |
| org.label-schema.url                     | https://github.com/jenkins-infra/docker-terraform.git          | https://github.com/jenkins-infra/docker-terraform.git          |
| org.opencontainers.image.revision        | ad902ec                                                        | 441c261                                                        |
| org.opencontainers.image.url             | https://github.com/jenkins-infra/docker-terraform.git          | https://github.com/jenkins-infra/docker-terraform.git          |
| io.jenkins-infra.tools.golang.version    | 1.15                                                           | 1.15                                                           |
+------------------------------------------+----------------------------------------------------------------+----------------------------------------------------------------+
| GitHub URL                               | https://github.com/jenkins-infra/docker-terraform/tree/ad902ec | https://github.com/jenkins-infra/docker-terraform/tree/441c261 |
+------------------------------------------+----------------------------------------------------------------+----------------------------------------------------------------+
https://github.com/jenkins-infra/docker-terraform/compare/ad902ec..441c261
```
