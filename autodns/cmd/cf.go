package cmd

import (
	"os"

	"github.com/libdns/cloudflare"
	"github.com/lolPants/autodns/autodns/pkg/constants"
	"github.com/lolPants/autodns/autodns/pkg/dnstools"
	"github.com/lolPants/autodns/autodns/pkg/iptools"
	"github.com/lolPants/autodns/autodns/pkg/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	cfTokenEnv = "AUTODNS_CF_TOKEN"
)

var (
	cfCmd = &cobra.Command{
		Use:   "cf",
		Short: "Use CloudFlare DNS provider",
		Run: func(cmd *cobra.Command, args []string) {
			record := viper.GetString("record")
			if record == "" {
				logger.Stderr.Println(0, constants.ErrorMissingRecord)
				os.Exit(1)
			}

			token := viper.GetString("cf-token")
			if token == "" {
				logger.Stderr.Println(0, constants.ErrorMissingToken(cfTokenEnv))
				os.Exit(1)
			}

			cf := &cloudflare.Provider{
				APIToken: token,
			}

			ip := <-iptools.Lookup()
			if ip.Error != nil {
				logger.Stderr.Println(0, constants.ErrorIPLookupFailed)
				os.Exit(1)
			}

			err := dnstools.SetRecords(cf, ip, record)
			if err != nil {
				logger.Stderr.Println(0, constants.ErrorRecordSetFailed)
				os.Exit(1)
			}
		},
	}
)

func init() {
	cfCmd.Flags().StringP("token", "T", "", "CloudFlare authentication token")
	viper.BindPFlag("cf-token", cfCmd.Flags().Lookup("token"))
	viper.BindEnv("cf-token", "AUTODNS_CF_TOKEN")

	cfCmd.Flags().StringP("record", "r", "", "DNS record")
	viper.BindPFlag("record", cfCmd.Flags().Lookup("record"))

	rootCmd.AddCommand(cfCmd)
}
