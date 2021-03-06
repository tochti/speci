package speci

import (
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func Test_SQLite(t *testing.T) {
	os.Clearenv()
	os.Setenv("T_SQLITE_PATH", ":memory:")

	sqlS, err := ReadSQLite("t")
	if err != nil {
		t.Fatal(err)
	}

	db, err := sqlS.DB()
	if err != nil {
		t.Fatal("Expect nil was %v", err)
	}

	err = db.Ping()
	if err != nil {
		t.Fatal(err)
	}

	sqlStr := "file::memory:"
	if sqlS.String() != sqlStr {
		t.Fatalf("Expect %v was %v", sqlStr, sqlS.String())
	}
}

func Test_WrongSQLiteConfig(t *testing.T) {
	os.Clearenv()

	_, err := ReadSQLite("t")
	if err == nil {
		t.Fatal("Expect error")
	}
}
