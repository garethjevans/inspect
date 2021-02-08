package inspect

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Client struct {
	Client http.Client
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

	tokenResponse := TokenResponse{}

	body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 10000000))
	if err != nil {
		return "", err
	}

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

	body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 10000000))
	if err != nil {
		return nil, err
	}

	manifestResponse := ManifestResponse{}
	err = json.Unmarshal(body, &manifestResponse)
	if err != nil {
		return nil, err
	}

	req, err = http.NewRequest("GET", fmt.Sprintf("https://registry-1.docker.io/v2/%s/blobs/%s", repo, manifestResponse.Config.Digest), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err = i.Client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err = ioutil.ReadAll(io.LimitReader(resp.Body, 10000000))
	if err != nil {
		return nil, err
	}

	blobResponse := BlobResponse{}
	err = json.Unmarshal(body, &blobResponse)
	if err != nil {
		return nil, err
	}

	return blobResponse.Config.Labels, nil
}
