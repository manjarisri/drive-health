package config

import (
	"os"
	"strconv"
)

type DHConfig struct {
	CleanupServiceFrequency int `json:"cleanupServiceFrequency"`
	DiskFetchFrequency      int `json:"diskFetchFrequency"`
	MaxHistoryAge           int `json:"maxHistoryAge"`

	DatabaseFilePath string `json:"databaseFilePath"`

	Listen string `json:"listen"`

	IdentityUsername string `json:"identityUsername"`
	IdentityPassword string `json:"identityPassword"`

	DebugMode bool `json:"debugMode"`
}

func GetConfiguration() DHConfig {
	config := DHConfig{
		DiskFetchFrequency:      5,
		CleanupServiceFrequency: 3600,
		MaxHistoryAge:           2592000,
		DatabaseFilePath:        "./data.sqlite",
		IdentityUsername:        "admin",
		IdentityPassword:        "admin",

		Listen: ":8080",
	}

	if val, exists := os.LookupEnv("DISK_FETCH_FREQUENCY"); exists {
		if intValue, err := strconv.Atoi(val); err == nil {
			config.DiskFetchFrequency = intValue
		}
	}

	if val, exists := os.LookupEnv("CLEANUP_SERVICE_FREQUENCY"); exists {
		if intValue, err := strconv.Atoi(val); err == nil {
			config.CleanupServiceFrequency = intValue
		}
	}

	if val, exists := os.LookupEnv("MAX_HISTORY_AGE"); exists {
		if intValue, err := strconv.Atoi(val); err == nil {
			config.MaxHistoryAge = intValue
		}
	}

	if val, exists := os.LookupEnv("LISTEN"); exists {
		config.Listen = val
	}

	if val, exists := os.LookupEnv("DATABASE_FILE_PATH"); exists {
		config.DatabaseFilePath = val
	}

	if val, exists := os.LookupEnv("IDENTITY_USERNAME"); exists {
		config.IdentityUsername = val
	}

	if val, exists := os.LookupEnv("IDENTITY_PASSWORD"); exists {
		config.IdentityPassword = val
	}

	if val, exists := os.LookupEnv("DEBUG_MODE"); exists {
		if isDebug, err := strconv.ParseBool(val); err == nil {
			config.DebugMode = isDebug
		}
	}

	return config
}
