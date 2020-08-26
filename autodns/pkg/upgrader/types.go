package upgrader

// Release represents a GitHub Release
type Release struct {
	ID         int64  `json:"id"`
	Tag        string `json:"tag_name"`
	Name       string `json:"name"`
	Draft      bool   `json:"draft"`
	Prerelease bool   `json:"prerelease"`
	URL        string `json:"url"`
	AssetsURL  string `json:"assets_url"`

	Assets []ReleaseAsset `json:"assets"`
}

// ReleaseAsset represents a GitHub Release Asset
type ReleaseAsset struct {
	ID          int64  `json:"id"`
	URL         string `json:"url"`
	Name        string `json:"name"`
	Label       string `json:"label"`
	DownloadURL string `json:"browser_download_url"`
}

// ReleaseResult represents an async result of type `Release`
type ReleaseResult struct {
	Error error
	Data  *Release
}
