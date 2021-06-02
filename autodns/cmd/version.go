package cmd

import (
	"fmt"
	"runtime"

	"github.com/lolPants/autodns/autodns/pkg/constants"
	"github.com/spf13/cobra"
)

var (
	sha1ver = constants.VersionUnknown
	gitTag  string
)

type versionRow struct {
	label string
	value string
}

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Run: func(cmd *cobra.Command, args []string) {
			printVersionInfo()
		},
	}
)

func printVersionInfo() {
	versionRows := make([]versionRow, 0)
	addRow := func(l, v string) {
		row := versionRow{label: l, value: v}
		versionRows = append(versionRows, row)
	}

	var version string
	if gitTag == "" {
		version = constants.VersionDev
	} else {
		version = gitTag
	}

	addRow("Version", version)
	addRow("Git Hash", sha1ver)
	addRow("Go Version", runtime.Version())

	var widest int
	for _, r := range versionRows {
		width := len(r.label) + 2
		if width > widest {
			widest = width
		}
	}

	for _, r := range versionRows {
		fmt.Printf("%*s %s\n", widest*-1, r.label, r.value)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
