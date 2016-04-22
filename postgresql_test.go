package speci

import (
	"os"
	"testing"

	_ "github.com/lib/pq"
)

func TestPostgreSQL(t *testing.T) {
	os.Clearenv()
	os.Setenv("T_POSTGRESQL_USER", "foo")
	os.Setenv("T_POSTGRESQL_PASSWORD", "bar")
	os.Setenv("T_POSTGRESQL_HOST", "h")
	os.Setenv("T_POSTGRESQL_PORT", "p")
	os.Setenv("T_POSTGRESQL_DB_NAME", "n")
	os.Setenv("T_POSTGRESQL_SSL_MODE", "allow")

	AppName = "t"
	sqlS, err := ReadPostgreSQL()
	if err != nil {
		t.Fatal(err)
	}

	_, err = sqlS.DB()
	if err != nil {
		t.Fatal("Expect nil was %v", err)
	}

	sqlStr := "postgres://foo:bar@h:p/n?sslmode=allow"
	if sqlS.String() != sqlStr {
		t.Fatalf("Expect %v was %v", sqlStr, sqlS.String())
	}
}

func TestPostgreSQL_MissingSSLMode(t *testing.T) {
	os.Clearenv()

	os.Setenv("T_POSTGRESQL_USER", "foo")
	os.Setenv("T_POSTGRESQL_PASSWORD", "bar")
	os.Setenv("T_POSTGRESQL_HOST", "h")
	os.Setenv("T_POSTGRESQL_PORT", "p")
	os.Setenv("T_POSTGRESQL_DB_NAME", "n")

	AppName = "t"
	sqlS, err := ReadPostgreSQL()

	if err != nil {
		t.Fatal(err)
	}

	sqlStr := "postgres://foo:bar@h:p/n?sslmode=disable"
	if sqlS.String() != sqlStr {
		t.Fatalf("Expect %v was %v", sqlStr, sqlS.String())
	}

}

func TestPostgreSQL_WrongConfig(t *testing.T) {
	os.Clearenv()

	AppName = "t"
	_, err := ReadPostgreSQL()
	if err == nil {
		t.Fatal("Expect error")
	}

}
