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

```
make test
```
