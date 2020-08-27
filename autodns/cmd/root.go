package cmd

import (
	"os"

	"github.com/lolPants/autodns/autodns/pkg/constants"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "autodns",
		Short: "A simple tool to automatically update DNS records for a server.",
		Long: "A simple tool to automatically update DNS records for a server.\n" +
			"More information is available at " + constants.GitHubURL,
		Run: func(cmd *cobra.Command, args []string) {
			if viper.GetBool("version") {
				versionCmd.Run(cmd, args)
				os.Exit(0)
			}

			cmd.Help()
			os.Exit(0)
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().CountP("verbose", "v", "verbose output")
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))

	rootCmd.Flags().BoolP("version", "V", false, "print version")
	viper.BindPFlag("version", rootCmd.Flags().Lookup("version"))
}
