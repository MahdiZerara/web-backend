package config

import (
	"fmt"
	"os"
)

// App server configs
type AppServerConfig struct {
	restPort string
}

func newAppServerConfig() (*AppServerConfig, error) {
	restPort := os.Getenv("REST_SERVER_PORT")
	if restPort == "" {
		return nil, fmt.Errorf("unable to retrieve 'REST_SERVER_PORT' from the environment variables")
	}

	return &AppServerConfig{
		restPort: restPort,
	}, nil
}

func (c *AppServerConfig) GetRESTPort() string {
	return c.restPort
}
