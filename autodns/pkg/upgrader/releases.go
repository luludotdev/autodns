package upgrader

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/lolPants/autodns/autodns/pkg/constants"
	"github.com/lolPants/autodns/autodns/pkg/logger"
)

// Fetch gets a GitHub release for a specific tag
func Fetch(tag string) (*Release, error) {
	url := "https://api.github.com/repos/" + constants.RepoSlug + "/releases/" + tag
	resp, err := client.Get(url)
	if err != nil {
		logger.Stderr.Println(1, "failed to fetch latest github release")
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		code := strconv.Itoa(resp.StatusCode)
		logger.Stderr.Println(1, "latest github release returned status code `"+code+"`")

		return nil, errors.New("status code not OK")
	}

	var release Release
	err = json.NewDecoder(resp.Body).Decode(&release)
	if err != nil {
		logger.Stderr.Println(1, "failed to deserialize release json")

		return nil, err
	}

	return &release, nil
}

// Latest gets a the latest GitHub release
func Latest() (*Release, error) {
	return Fetch("latest")
}

// FetchAsync asynchronously gets a GitHub release for a specific tag
func FetchAsync(tag string) <-chan *ReleaseResult {
	r := make(chan *ReleaseResult)

	go func() {
		defer close(r)
		res := &ReleaseResult{}

		release, err := Fetch(tag)
		if err != nil {
			res.Error = err
		} else {
			res.Data = release
		}

		r <- res
	}()

	return r
}

// LatestAsync asynchronously gets a the latest GitHub release
func LatestAsync() <-chan *ReleaseResult {
	return FetchAsync("latest")
}
