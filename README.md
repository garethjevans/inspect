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
