package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/bobesa/go-domain-util/domainutil"
	"github.com/libdns/cloudflare"
	"github.com/libdns/libdns"
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

	records := make([]libdns.Record, 0)

	if v4ptr != nil {
		r := libdns.Record{
			Type:  "A",
			Name:  record,
			Value: *v4ptr,
		}

		records = append(records, r)
	}

	if v6ptr != nil {
		r := libdns.Record{
			Type:  "AAAA",
			Name:  record,
			Value: *v6ptr,
		}

		records = append(records, r)
	}

	ctx := context.Background()
	domain := domainutil.Domain(record)

	_, err := cf.SetRecords(ctx, domain, records)
	if err != nil {
		fmt.Println("Failed to set DNS Records!")
		fmt.Println(err)

		os.Exit(1)
	}
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
