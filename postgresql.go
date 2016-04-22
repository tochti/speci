package speci

import (
	"database/sql"
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type (
	PostgreSQL struct {
		User     string `envconfig:"POSTGRESQL_USER" required:"true"`
		Password string `envconfig:"POSTGRESQL_PASSWORD" required:"true"`
		Host     string `envconfig:"POSTGRESQL_HOST" required:"true"`
		Port     string `envconfig:"POSTGRESQL_PORT" required:"true"`
		DBName   string `envconfig:"POSTGRESQL_DB_NAME" required:"true"`
		SSLMode  string `envconfig:"POSTGRESQL_SSL_MODE"`
	}
)

// Read PostgreSQL settings from env
func ReadPostgreSQL() (*PostgreSQL, error) {
	sqlS := &PostgreSQL{}

	err := envconfig.Process(AppName, sqlS)

	if err != nil {
		return nil, err
	}

	return sqlS, nil
}

func (m *PostgreSQL) DB() (*sql.DB, error) {
	db, err := sql.Open("postgres", m.String())
	if err != nil {
		return db, err
	}

	return db, nil
}

func (m *PostgreSQL) String() string {
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		m.User,
		m.Password,
		m.Host,
		m.Port,
		m.DBName,
	)

	if m.SSLMode == "" {
		url = fmt.Sprintf("%v?sslmode=disable", url)
	} else {
		url = fmt.Sprintf("%v?sslmode=%v", url, m.SSLMode)
	}

	return url
}
