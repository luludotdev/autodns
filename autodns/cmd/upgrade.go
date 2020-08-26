package cmd

import (
	"os"

	"github.com/lolPants/autodns/autodns/pkg/constants"
	"github.com/lolPants/autodns/autodns/pkg/logger"
	"github.com/lolPants/autodns/autodns/pkg/upgrader"
	"github.com/spf13/cobra"
)

var (
	upgradeCommand = &cobra.Command{
		Use:   "upgrade",
		Short: "Upgrade autodns to the latest version",
		Run: func(cmd *cobra.Command, args []string) {
			if upgrader.IsDev(gitTag) {
				logger.Stdout.Println(0, "You are running a dev build!")
				logger.Stdout.Println(0, "Visit "+constants.GitHubURL+"/releases to download the latest release.")

				os.Exit(0)
			}

			release, err := upgrader.Latest()
			if err != nil {
				logger.Stderr.Printf(1, "failed to get latest release, error: `%s`\n", err.Error())
				logger.Stderr.Println(0, "Failed to check for new versions!")
				logger.Stderr.Println(0, constants.NewReleasesMessage)

				os.Exit(1)
			}

			needsUpgrade, err := upgrader.NeedsUpgrade(gitTag, release.Tag)
			if err != nil {
				logger.Stderr.Printf(1, "failed to coerce versions to semver, error: `%s`\n", err.Error())
				logger.Stderr.Println(0, "Failed to check for new versions!")
				logger.Stderr.Println(0, constants.NewReleasesMessage)

				os.Exit(1)
			}

			if needsUpgrade == false {
				logger.Stdout.Println(0, "You are running the latest version!")
				os.Exit(0)
			}

			asset := release.GetAsset()
			if asset == nil {
				logger.Stderr.Println(0, "Failed to resolve download for your platform!")
				logger.Stderr.Println(0, constants.NewReleasesMessage)

				os.Exit(1)
			}

			reader, err := asset.Download()
			if err != nil {
				logger.Stderr.Printf(1, "asset download failed, error: `%s`\n", err.Error())
				logger.Stderr.Println(0, "Failed to download latest version!")
				logger.Stderr.Println(0, constants.NewReleasesMessage)

				os.Exit(1)
			}

			defer reader.Close()
			err = upgrader.Replace(reader)
			if err != nil {
				logger.Stderr.Printf(1, "binary replace failed, error: `%s`\n", err.Error())
				logger.Stderr.Println(0, "Failed to replace current binary!")
				logger.Stderr.Println(0, constants.NewReleasesMessage)

				os.Exit(1)
			}

			logger.Stdout.Printf(0, "Upgraded %s to %s\n", constants.Name, release.Tag)
		},
	}
)

func init() {
	_ = upgrader.Cleanup()

	rootCmd.AddCommand(upgradeCommand)
}
