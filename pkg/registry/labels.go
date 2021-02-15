package registry

import (
	"fmt"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"strings"
)

type LabelLister interface {
	Labels(repo string, version string) (map[string]string, error)
}

type DefaultLabelLister struct {

}

func (d *DefaultLabelLister) Labels(repo string, version string) (map[string]string, error) {
	var imageImage string
	if strings.HasPrefix(version, "sha256:") {
		imageImage = fmt.Sprintf("%s@%s", repo, version)
	} else {
		imageImage = fmt.Sprintf("%s:%s", repo, version)
	}
	ref, err := name.ParseReference(imageImage)
	if err != nil {
		return nil, err
	}

	img, err := remote.Image(ref)
	if err != nil {
		return nil, err
	}

	configFile, err := img.ConfigFile()
	if err != nil {
		return nil, err
	}

	return configFile.Config.Labels, nil
}
