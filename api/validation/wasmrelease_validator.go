package validation

import (
	"regexp"

	"github.com/boeboe/wasm-repo/api/errors"
	"github.com/boeboe/wasm-repo/api/models"
)

var (
	validVersionRegex = regexp.MustCompile(`^[a-zA-Z0-9-_.]+$`)
)

func ValidateWASMRelease(release *models.WASMRelease) error {
	if err := isValidVersion(release.Version); err != nil {
		return err
	}
	return nil
}

func isValidVersion(version string) error {
	if !validVersionRegex.MatchString(version) {
		return &errors.ValidationError{
			Field:   "Version",
			Message: "Invalid release version. It should only contain alphanumeric characters, dashes, underscores, and dots.",
		}
	}
	return nil
}
