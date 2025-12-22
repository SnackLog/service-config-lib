package serviceconfiglib

import "os"

var activeConfig ServiceConfig

// LoadConfig Load configuration from environment variables and validate it
func LoadConfig() {
	activeConfig.AppName = os.Getenv("SERVICE_CONFIG_APP_NAME")
	activeConfig.ApiRootUrl = os.Getenv("SERVICE_CONFIG_API_ROOT_URL")
	activeConfig.ServiceName = os.Getenv("SERVICE_CONFIG_SERVICE_NAME")
}

// SetConfig Set the active configuration
func SetConfig(config ServiceConfig) {
	activeConfig = config
}

// GetConfig Return a copy of the active configuration
func GetConfig() ServiceConfig {
	return activeConfig
}
