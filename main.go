package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/libdns/cloudflare"
	"github.com/lolPants/flaggs"
)

var (
	apiToken = ""
	record   = ""

	cf = &cloudflare.Provider{}
)

func main() {
	flaggs.SetDetails("AutoDNS "+gitTag, "https://github.com/lolPants/autodns")

	flaggs.RegisterStringFlag(&apiToken, "T", "api-token", "CloudFlare API Token")
	flaggs.RegisterStringFlag(&record, "r", "record", "DNS Record")
	flaggs.RegisterBoolFlag(&printVersion, "v", "version", "Print version information")
	flaggs.Parse(nil)

	if printVersion == true {
		printVersionInfo()
		os.Exit(0)
	}

	envToken := os.Getenv("CLOUDFLARE_TOKEN")
	if envToken != "" {
		apiToken = envToken
	}

	if apiToken == "" {
		fmt.Println("Missing API token! Must be specified with --api-token flag or CLOUDFLARE_TOKEN environment variable.")
		os.Exit(1)
	}

	if record == "" {
		fmt.Println("Missing DNS Record! Must be specified with --record flag.")
		os.Exit(1)
	}

	cf.APIToken = apiToken

	v4task, v6task := getAddress("v4"), getAddress("v6")
	v4ptr, v6ptr := <-v4task, <-v6task

	fmt.Printf("%+v\t%+v\n", v4ptr, v6ptr)
}

func getAddress(subdomain string) <-chan *string {
	r := make(chan *string)

	go func() {
		defer close(r)

		var client = &http.Client{
			Timeout: time.Second * 10,
		}

		url := "https://" + subdomain + ".ipv6-test.com/api/myip.php"
		resp, err := client.Get(url)
		if err != nil {
			r <- nil
			return
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			r <- nil
			return
		}

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			r <- nil
			return
		}

		str := string(bodyBytes)
		r <- &str
	}()

	return r
}
