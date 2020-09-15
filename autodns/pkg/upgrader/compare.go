package upgrader

import (
	"github.com/Masterminds/semver/v3"
	"github.com/lolPants/autodns/autodns/pkg/constants"
	"github.com/lolPants/autodns/autodns/pkg/logger"
	"github.com/spf13/viper"
)

// IsDev checks if a given version is in dev mode
func IsDev(version string) bool {
	logger.Stdout.Println(2, "checking if current version is a dev build")
	if version == "" {
		return true
	}

	if version == constants.VersionDev {
		return true
	}

	return false
}

// NeedsUpgrade calculates if the current version needs upgrading
func NeedsUpgrade(version string, latest string) (bool, error) {
	logger.Stdout.Println(2, "checking if upgrade is needed")

	force := viper.GetBool("force-upgrade")
	if IsDev(version) {
		if force {
			logger.Stdout.Println(2, "is dev, forcing upgrade")
			return true, nil
		}

		logger.Stdout.Println(2, "is dev, upgrade not needed")
		return false, nil
	}

	latestVer, err := semver.NewVersion(latest)
	if err != nil {
		logger.Stderr.Printf(1, "failed to parse latest version, error: `%s`\n", err.Error())
		return false, err
	}

	currentVer, err := semver.NewVersion(version)
	if err != nil {
		logger.Stderr.Printf(1, "failed to parse current version, error: `%s`\n", err.Error())
		return false, err
	}

	return latestVer.GreaterThan(currentVer), nil
}
