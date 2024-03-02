package config

import (
	"fmt"
)

// Global App Configs
type AppConfig struct {
	SharedConfig     *SharedConfig
	PostgresDBConfig *PostgresDBConfig
	AppServerConfig  *AppServerConfig
}

func NewAppConfig() (*AppConfig, error) {
	sharedConfig, newSharedConfigErr := newSharedConfig()
	if newSharedConfigErr != nil {
		return nil, fmt.Errorf("an error occurred while retrieving the shared configs: %v", newSharedConfigErr)
	}

	postgresDBConfig, newPostgresDBConfigErr := newPostgresDBConfig()
	if newPostgresDBConfigErr != nil {
		return nil, fmt.Errorf("an error occurred while retrieving the postgres database configs: %v", newPostgresDBConfigErr)
	}

	appServerConfig, newAppServerConfigErr := newAppServerConfig()
	if newAppServerConfigErr != nil {
		return nil, fmt.Errorf("an error occurred while retrieving the app server configs: %v", newAppServerConfigErr)
	}

	return &AppConfig{
		SharedConfig:     sharedConfig,
		PostgresDBConfig: postgresDBConfig,
		AppServerConfig:  appServerConfig,
	}, nil
}
