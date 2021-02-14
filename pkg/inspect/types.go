package inspect

import (
	"fmt"
	"strings"
)

// TokenResponse struct containing the response of a token request.
type TokenResponse struct {
	Token string `json:"token"`
}

// Config struct containing the config response.
type Config struct {
	MediaType string `json:"mediaType"`
	Size      int    `json:"size"`
	Digest    string `json:"digest"`
}

// ManifestResponse a struct containing all the details for a manifest response.
type ManifestResponse struct {
	SchemaVersion int      `json:"schemaVersion"`
	MediaType     string   `json:"mediaType"`
	Config        Config   `json:"config"`
	Layers        []Config `json:"layers"`
}

// BlobResponse a struct containing all the details for a blob response.
type BlobResponse struct {
	Architecture string     `json:"architecture"`
	Config       BlobDetail `json:"config"`
}

// BlobDetail the details of a blob layer.
type BlobDetail struct {
	Cmd          []string `json:"Cmd"`
	Entrypoint   []string `json:"Entrypoint"`
	Env          []string `json:Env`
	Hostname     string   `json:"Hostname"`
	DomainName   string   `json:"Domainname"`
	User         string   `json:"User"`
	Image        string   `json:"Image"`
	WorkingDir   string   `json:"WorkingDir"`
	OnBuild      string   `json:"OnBuild"`
	AttachStdin  bool     `json:"AttachStdin"`
	AttachStdout bool     `json:"AttachStdout"`
	AttachStderr bool     `json:"AttachStderr"`
	OpenStdin    bool     `json:"OpenStdin"`
	StdinOnce    bool     `json:"StdinOnce"`
	Tty          bool     `json:"Tty"`
	// Volumes
	//"ExposedPorts":{"389/tcp":{},"636/tcp”:{}}
	Labels map[string]string `json:"Labels"`
}

// SourceURL Gets the SourceURL for the revision.
func SourceURL(labels map[string]string) string {
	return first(labels, "org.opencontainers.image.url", "org.label-schema.url")
}

// GitHubURL Gets the GitHubUrl for the revision.
func GitHubURL(labels map[string]string) string {
	return fmt.Sprintf("%s/tree/%s", BaseURL(labels), Revision(labels))
}

// BaseURL Gets the base source url without the .git suffix.
func BaseURL(labels map[string]string) string {
	gitURL := SourceURL(labels)
	if strings.HasSuffix(gitURL, ".git") {
		gitURL = strings.TrimSuffix(gitURL, ".git")
	}
	return gitURL
}

// Revision get the commit revision for this image.
func Revision(labels map[string]string) string {
	return first(labels, "org.opencontainers.image.revision", "org.label-schema.vcs-ref")
}

func first(labels map[string]string, names ...string) string {
	for _, n := range names {
		r := labels[n]
		if r != "" {
			return r
		}
	}
	return ""
}
