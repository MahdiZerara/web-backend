package config

import (
	"fmt"
	"os"
)

// Postgres database configs
type PostgresDBConfig struct {
	host     string
	name     string
	user     string
	password string
	port     string
}

func newPostgresDBConfig() (*PostgresDBConfig, error) {
	host := os.Getenv("DATABASE_HOST")
	if host == "" {
		return nil, fmt.Errorf("unable to retrieve 'DATABASE_HOST' from the environment variables")
	}

	name := os.Getenv("DATABASE_NAME")
	if name == "" {
		return nil, fmt.Errorf("unable to retrieve 'DATABASE_NAME' from the environment variables")
	}

	user := os.Getenv("DATABASE_USER")
	if user == "" {
		return nil, fmt.Errorf("unable to retrieve 'DATABASE_USER' from the environment variables")
	}

	password := os.Getenv("DATABASE_PASSWORD")
	if password == "" {
		return nil, fmt.Errorf("unable to retrieve 'DATABASE_PASSWORD' from the environment variables")
	}

	port := os.Getenv("DATABASE_PORT")
	if port == "" {
		return nil, fmt.Errorf("unable to retrieve 'DATABASE_PORT' from the environment variables")
	}

	return &PostgresDBConfig{
		host:     host,
		name:     name,
		user:     user,
		password: password,
		port:     port,
	}, nil
}

func (c *PostgresDBConfig) GetDBHost() string {
	return c.host
}

func (c *PostgresDBConfig) GetDBName() string {
	return c.name
}

func (c *PostgresDBConfig) GetDBUser() string {
	return c.user
}

func (c *PostgresDBConfig) GetDBPassword() string {
	return c.password
}

func (c *PostgresDBConfig) GetDBPort() string {
	return c.port
}
