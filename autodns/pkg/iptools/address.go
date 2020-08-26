package iptools

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/lolPants/autodns/autodns/pkg/logger"
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

		logger.Stdout.Println(2, "resolving IP"+subdomain+" address")
		url := "https://" + subdomain + ".ipv6-test.com/api/myip.php"
		resp, err := client.Get(url)

		if err != nil {
			logger.Stderr.Println(1, "failed to resolve IP"+subdomain+" address!")
			if strings.Contains(err.Error(), "no such host") == false {
				res.Error = err
			}

			r <- res
			return
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			code := strconv.Itoa(resp.StatusCode)
			logger.Stderr.Println(1, "IP"+subdomain+" resolver returned status code `"+code+"`")

			res.Error = errors.New("status code not OK")
			r <- res

			return
		}

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logger.Stderr.Println(1, "failed to read IP"+subdomain+" body!")
			res.Error = err
			r <- res

			return
		}

		str := string(bodyBytes)
		logger.Stdout.Println(1, "resolved IP"+subdomain+" address as `"+str+"`")
		res.Data = &str

		r <- res
	}()

	return r
}
