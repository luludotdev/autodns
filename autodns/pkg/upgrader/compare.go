package upgrader

import (
	"strings"

	"github.com/Masterminds/semver"
	"github.com/lolPants/autodns/autodns/pkg/constants"
)

// IsDev checks if a given version is in dev mode
func IsDev(version string) bool {
	if version == "" {
		return true
	}

	if version == constants.VersionDev {
		return true
	}

	return false
}

// NeedsUpgrade calculates if the current version needs upgrading
func NeedsUpgrade(version string) (bool, error) {
	if IsDev(version) {
		return false, nil
	}

	release, err := Latest()
	if err != nil {
		return false, err
	}

	latest, err := semver.NewVersion(strings.TrimLeft(release.Tag, "v"))
	if err != nil {
		return false, err
	}

	current, err := semver.NewVersion(strings.TrimLeft(version, "v"))
	if err != nil {
		return false, err
	}

	return latest.GreaterThan(current), nil
}
