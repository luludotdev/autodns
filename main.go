package main

import (
	"fmt"
	"os"

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
}
