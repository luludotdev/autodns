package upgrader

import (
	"strings"

	"github.com/Masterminds/semver"
	"github.com/lolPants/autodns/autodns/pkg/constants"
)

// IsDev checks if a given version is in dev mode
func IsDev(version string) bool {
	return version == constants.VersionDev
}

// NeedsUpgrade calculates if the current version needs upgrading
func NeedsUpgrade(version string) (bool, error) {
	release, err := Latest()
	if err != nil {
		return false, err
	}

	tag := strings.TrimLeft(release.Tag, "v")
	latest, err := semver.NewVersion(tag)
	if err != nil {
		return false, err
	}

	current, err := semver.NewVersion(version)
	if err != nil {
		return false, err
	}

	return latest.GreaterThan(current), nil
}
