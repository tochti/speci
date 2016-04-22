package speci

import (
	"database/sql"
	"errors"
	"fmt"
	"net/url"

	"github.com/kelseyhightower/envconfig"
)

const (
	DefaultLocation = "Europe/Berlin"
)

var (
	ErrMySQLConfig = errors.New("Wrong MySQL config")
)

type (
	MySQL struct {
		User     string `envconfig:"MYSQL_USER"`
		Password string `envconfig:"MYSQL_PASSWORD"`
		Host     string `envconfig:"MYSQL_HOST"`
		Port     string `envconfig:"MYSQL_PORT"`
		DBName   string `envconfig:"MYSQL_DB_NAME"`
		Location string `envconfig:"MYSQL_LOCATION"`
	}
)

// Read MySQL settings from env
func ReadMySQL() (*MySQL, error) {
	mysql := &MySQL{}

	err := envconfig.Process(AppName, mysql)

	if err != nil {
		return nil, nil
	}

	if mysql.Host == "" {
		return nil, ErrMySQLConfig
	}

	return mysql, nil
}

func (m *MySQL) DB() (*sql.DB, error) {
	db, err := sql.Open("mysql", m.String())
	if err != nil {
		return db, err
	}

	return db, nil
}

func (m *MySQL) String() string {
	if m.Location == "" {
		m.Location = url.QueryEscape(DefaultLocation)
	} else {
		m.Location = url.QueryEscape(m.Location)
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?loc=%s&parseTime=true",
		m.User,
		m.Password,
		m.Host,
		m.Port,
		m.DBName,
		m.Location,
	)
}
