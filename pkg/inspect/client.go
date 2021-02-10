package inspect

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

// HTTPClient interface that wraps the Do function.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	Client HTTPClient
}

func (i *Client) token(repo string) (string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://auth.docker.io/token?service=registry.docker.io&scope=repository:%s:pull", repo), nil)
	if err != nil {
		return "", err
	}

	resp, err := i.Client.Do(req)
	if err != nil {
		return "", err
	}
	logrus.Debugf("got status code %d", resp.StatusCode)

	tokenResponse := TokenResponse{}

	body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 10000000))
	if err != nil {
		return "", err
	}

	logrus.Debugf("body> %s", string(body))

	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		return "", err
	}

	return tokenResponse.Token, nil
}

func (i *Client) Labels(repo string, version string) (map[string]string, error) {
	token, err := i.token(repo)
	if err != nil {
		return nil, err
	}

	var digest string

	req, err := http.NewRequest("GET", fmt.Sprintf("https://registry-1.docker.io/v2/%s/manifests/%s", repo, version), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Accept", "application/vnd.docker.distribution.manifest.v2+json")

	resp, err := i.Client.Do(req)
	if err != nil {
		return nil, err
	}

	logrus.Debugf("got status code %d", resp.StatusCode)

	if resp.StatusCode != 200 {
		return nil, errors.New("unable to find manifest " + version + ", " + resp.Status)
	}

	body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 10000000))
	if err != nil {
		return nil, err
	}

	logrus.Debugf("body> %s", string(body))

	manifestResponse := ManifestResponse{}
	err = json.Unmarshal(body, &manifestResponse)
	if err != nil {
		return nil, err
	}

	logrus.Debugf("got digest %s", manifestResponse.Config.Digest)
	digest = manifestResponse.Config.Digest

	req, err = http.NewRequest("GET", fmt.Sprintf("https://registry-1.docker.io/v2/%s/blobs/%s", repo, digest), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err = i.Client.Do(req)
	if err != nil {
		return nil, err
	}
	logrus.Debugf("got status code %d", resp.StatusCode)
	if resp.StatusCode != 200 {
		return nil, errors.New("unable to find blob " + digest + ", " + resp.Status)
	}

	body, err = ioutil.ReadAll(io.LimitReader(resp.Body, 10000000))
	if err != nil {
		return nil, err
	}

	logrus.Debugf("body> %s", string(body))

	blobResponse := BlobResponse{}
	err = json.Unmarshal(body, &blobResponse)
	if err != nil {
		return nil, err
	}

	return blobResponse.Config.Labels, nil
}
