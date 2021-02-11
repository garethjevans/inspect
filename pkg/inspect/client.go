package inspect

import (
	"encoding/json"
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

func (i *Client) get(url string, response interface{}, headers http.Header) error {
	logrus.Debugf("requesting %s", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	if headers != nil {
		req.Header = headers
	}

	resp, err := i.Client.Do(req)
	if err != nil {
		return err
	}

	logrus.Debugf("got status code %d", resp.StatusCode)

	if resp.StatusCode != 200 {
		return fmt.Errorf("unexpected response status %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 10000000))
	if err != nil {
		return err
	}

	logrus.Debugf("body> %s", string(body))
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	return nil
}

func (i *Client) token(repo string) (string, error) {
	tokenResponse := TokenResponse{}

	err := i.get(fmt.Sprintf("https://auth.docker.io/token?service=registry.docker.io&scope=repository:%s:pull", repo), &tokenResponse, nil)
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

	manifestResponse := ManifestResponse{}

	headers := http.Header{}
	headers.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	headers.Add("Accept", "application/vnd.docker.distribution.manifest.v2+json")

	err = i.get(fmt.Sprintf("https://registry-1.docker.io/v2/%s/manifests/%s", repo, version), &manifestResponse, headers)
	if err != nil {
		return nil, err
	}

	logrus.Debugf("got digest %s", manifestResponse.Config.Digest)
	digest = manifestResponse.Config.Digest

	headers = http.Header{}
	headers.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	blobResponse := BlobResponse{}

	err = i.get(fmt.Sprintf("https://registry-1.docker.io/v2/%s/blobs/%s", repo, digest), &blobResponse, headers)
	if err != nil {
		return nil, err
	}

	return blobResponse.Config.Labels, nil
}
