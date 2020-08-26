package cmd

import (
	"fmt"
	"os"

	"github.com/libdns/cloudflare"
	"github.com/lolPants/autodns/autodns/pkg/constants"
	"github.com/lolPants/autodns/autodns/pkg/dnstools"
	"github.com/lolPants/autodns/autodns/pkg/iptools"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	cfTokenEnv = "AUTODNS_CF_TOKEN"
)

var (
	cfCommand = &cobra.Command{
		Use:   "cf",
		Short: "Use CloudFlare DNS provider",
		Run: func(cmd *cobra.Command, args []string) {
			record := viper.GetString("record")
			if record == "" {
				fmt.Println(constants.ErrorMissingRecord)
				os.Exit(1)
			}

			token := viper.GetString("cf-token")
			if token == "" {
				fmt.Println(constants.ErrorMissingToken(cfTokenEnv))
				os.Exit(1)
			}

			cf := &cloudflare.Provider{
				APIToken: token,
			}

			ip := <-iptools.Lookup()
			if ip.Error != nil {
				fmt.Println(constants.ErrorIPLookupFailed)
				os.Exit(1)
			}

			err := dnstools.SetRecords(cf, ip, record)
			if err != nil {
				fmt.Println(constants.ErrorRecordSetFailed)
				os.Exit(1)
			}
		},
	}
)

func init() {
	cfCommand.Flags().StringP("token", "T", "", "CloudFlare authentication token")
	viper.BindPFlag("cf-token", cfCommand.Flags().Lookup("token"))
	viper.BindEnv("cf-token", "AUTODNS_CF_TOKEN")

	cfCommand.Flags().StringP("record", "r", "", "DNS record")
	viper.BindPFlag("record", cfCommand.Flags().Lookup("record"))

	rootCmd.AddCommand(cfCommand)
}
