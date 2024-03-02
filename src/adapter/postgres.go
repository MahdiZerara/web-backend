package adapter

import (
	"database/sql"
	"fmt"

	"github.com/MahdiZerara/web-backend/config"
	_ "github.com/lib/pq"
)

type PostgresClient struct {
	SQLDB  *sql.DB
	config config.PostgresDBConfig
}

func NewPostgresClient(cfg *config.PostgresDBConfig) *PostgresClient {
	return &PostgresClient{
		config: *cfg,
	}
}

func (c *PostgresClient) Connect() error {
	connString := c.getDBConnectionString()
	db, openErr := sql.Open("postgres", connString)
	if openErr != nil {
		return openErr
	}

	c.SQLDB = db

	if err := c.Ping(); err != nil {
		return err
	}

	return nil
}

func (c *PostgresClient) Ping() error {
	return c.SQLDB.Ping()
}

func (c *PostgresClient) Disconnect() error {
	return c.SQLDB.Close()
}

func (c *PostgresClient) Create(sqlStatement string, args ...any) error {
	_, err := c.SQLDB.Exec(sqlStatement, args...)
	return err
}

func (c *PostgresClient) Read(sqlStatement string, args ...any) (*sql.Rows, error) {
	rows, queryErr := c.SQLDB.Query(sqlStatement, args...)
	if queryErr != nil {
		return nil, queryErr
	}

	return rows, nil
}

func (c *PostgresClient) getDBConnectionString() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.config.GetDBHost(),
		c.config.GetDBPort(),
		c.config.GetDBUser(),
		c.config.GetDBPassword(),
		c.config.GetDBName(),
	)
}
