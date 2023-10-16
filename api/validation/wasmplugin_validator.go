package validation

import (
	"regexp"
	"strings"

	"github.com/boeboe/wasm-repo/api/errors"
	"github.com/boeboe/wasm-repo/api/models"
)

const (
	HttpFilter    = "httpfilter"
	NetworkFilter = "networkfilter"
	WasmService   = "wasmservice"
)

var (
	validNameRegex = regexp.MustCompile(`^[a-zA-Z0-9-_]+$`)
	validTypes     = map[string]bool{
		HttpFilter:    true,
		NetworkFilter: true,
		WasmService:   true,
	}
)

func ValidateWASMPlugin(plugin *models.WASMPlugin) error {
	if err := isValidPluginName(plugin.Name); err != nil {
		return err
	}
	if err := isValidPluginType(plugin.Type); err != nil {
		return err
	}
	return nil
}

func isValidPluginName(name string) error {
	if !validNameRegex.MatchString(name) {
		return &errors.ValidationError{
			Field:   "Name",
			Message: "Invalid plugin name. It should only contain alphanumeric characters, dashes, and underscores.",
		}
	}
	return nil
}

func isValidPluginType(value string) error {
	if _, ok := validTypes[strings.ToLower(value)]; !ok {
		return &errors.ValidationError{
			Field:   "Type",
			Message: "Invalid plugin type. It should be one of httpfilter, networkfilter, or wasmservice.",
		}
	}
	return nil
}
