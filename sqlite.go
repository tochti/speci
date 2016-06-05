package speci

import (
	"database/sql"
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type (
	SQLite struct {
		Path string `envconfig:"SQLITE_PATH" required:"true"`
	}
)

func ReadSQLite(prefix string) (*SQLite, error) {
	specs := &SQLite{}

	err := envconfig.Process(prefix, specs)
	if err != nil {
		return &SQLite{}, err
	}

	return specs, nil
}

func (s *SQLite) DB() (*sql.DB, error) {
	pool, err := sql.Open("sqlite3", s.String())
	if err != nil {
		return &sql.DB{}, err
	}

	return pool, nil
}

func (s *SQLite) String() string {
	return fmt.Sprintf("file:%v", s.Path)
}
