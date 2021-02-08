package inspect

type TokenResponse struct {
	Token string `json:"token"`
}

type Config struct {
	MediaType string `json:"mediaType"`
	Size      int    `json:"size"`
	Digest    string `json:"digest"`
}

type ManifestResponse struct {
	SchemaVersion int      `json:"schemaVersion"`
	MediaType     string   `json:"mediaType"`
	Config        Config   `json:"config"`
	Layers        []Config `json:"layers"`
}

type BlobResponse struct {
	Architecture string     `json:"architecture"`
	Config       BlobDetail `json:"config"`
}

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
