package upgrader

import (
	"strings"

	"github.com/Masterminds/semver"
	"github.com/lolPants/autodns/autodns/pkg/constants"
	"github.com/lolPants/autodns/autodns/pkg/logger"
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
	if IsDev(version) {
		logger.Stdout.Println(2, "is dev, upgrade not needed")
		return false, nil
	}

	latestVer, err := semver.NewVersion(strings.TrimLeft(latest, "v"))
	if err != nil {
		logger.Stderr.Printf(1, "failed to parse latest version, error: `%s`\n", err.Error())
		return false, err
	}

	currentVer, err := semver.NewVersion(strings.TrimLeft(version, "v"))
	if err != nil {
		logger.Stderr.Printf(1, "failed to parse current version, error: `%s`\n", err.Error())
		return false, err
	}

	return latestVer.GreaterThan(currentVer), nil
}
