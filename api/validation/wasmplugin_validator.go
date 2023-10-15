package validation

import (
	"regexp"

	"github.com/boeboe/wasm-repo/api/errors"
)

// IsValidPluginName checks if the plugin name is valid.
func IsValidPluginName(name string) error {
	validName := regexp.MustCompile(`^[a-zA-Z0-9-_]+$`)
	if !validName.MatchString(name) {
		return &errors.ValidationError{
			Field:   "Name",
			Message: "Invalid plugin name. It should only contain alphanumeric characters, dashes, and underscores.",
		}
	}
	return nil
}
