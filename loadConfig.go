package serviceconfiglib

import (
	"fmt"
	"os"
)

var activeConfig ServiceConfig

// LoadConfig Load configuration from environment variables and validate it
func LoadConfig() error {
	var loadedConfig ServiceConfig
	loadedConfig.AppName = os.Getenv("SERVICE_CONFIG_APP_NAME")
	loadedConfig.ApiRootUrl = os.Getenv("SERVICE_CONFIG_API_ROOT_URL")
	loadedConfig.ServiceName = os.Getenv("SERVICE_CONFIG_SERVICE_NAME")

	err := ValidateConfig(loadedConfig)
	if err != nil {
		return fmt.Errorf("Config validation failed: %v", err)
	}
	return nil
}

// SetConfig Set the active configuration
func SetConfig(config ServiceConfig) {
	activeConfig = config
}

// GetConfig Return a copy of the active configuration
func GetConfig() ServiceConfig {
	return activeConfig
}
