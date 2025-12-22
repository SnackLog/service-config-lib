package serviceconfiglib

import "errors"

// ValidateConfig Validate the given ServiceConfig
func ValidateConfig(config ServiceConfig) error {
	if config.ServiceName == "" {
		return errors.New("Missing required field: ServiceName")
	}
	if config.ApiRootUrl == "" {
		return errors.New("Missing required field: ApiRootUrl")
	}

	return nil
}
