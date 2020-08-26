package iptools

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var (
	client = &http.Client{
		Timeout: time.Second * 10,
	}
)

type result struct {
	Error error
	Data  *string
}

func getAddress(subdomain string) <-chan *result {
	r := make(chan *result)

	go func() {
		defer close(r)
		res := &result{}

		url := "https://" + subdomain + ".ipv6-test.com/api/myip.php"
		resp, err := client.Get(url)
		if err != nil {
			if strings.Contains(err.Error(), "no such host") == false {
				res.Error = err
			}

			r <- res
			return
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			res.Error = errors.New("status code not OK")
			r <- res

			return
		}

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			res.Error = err
			r <- res

			return
		}

		str := string(bodyBytes)
		res.Data = &str

		r <- res
	}()

	return r
}
