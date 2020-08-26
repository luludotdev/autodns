package upgrader

import (
	"errors"
	"io"
	"net/http"
	"runtime"
	"strconv"
	"strings"

	"github.com/lolPants/autodns/autodns/pkg/logger"
)

// GetAsset gets the release asset for the current platform
func (r *Release) GetAsset() *ReleaseAsset {
	logger.Stdout.Println(2, "resolving release asset for the current platform")
	if len(r.Assets) == 0 {
		return nil
	}

	for _, asset := range r.Assets {
		if runtime.GOOS == "windows" && strings.Contains(asset.Name, ".exe") {
			return &asset
		}

		if runtime.GOOS == "linux" && (strings.Contains(asset.Name, ".exe") == false) {
			return &asset
		}
	}

	return nil
}

// Download gets the asset file as an io.ReadCloser
func (a *ReleaseAsset) Download() (io.ReadCloser, error) {
	logger.Stdout.Println(2, "sending http request for release asset: `"+a.Name+"`")
	resp, err := client.Get(a.DownloadURL)
	if err != nil {
		logger.Stderr.Println(1, "failed to fetch latest github release")
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		code := strconv.Itoa(resp.StatusCode)
		logger.Stderr.Println(1, "release asset status code `"+code+"`")

		return nil, errors.New("status code not OK")
	}

	return resp.Body, nil
}
