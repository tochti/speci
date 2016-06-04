package speci

import (
	"os"
	"testing"
)

func Test_SQLite(t *testing.T) {
	os.Clearenv()
	os.Setenv("T_SQLITE_PATH", ":memory:")

	sqlS, err := ReadSQLite("t")
	if err != nil {
		t.Fatal(err)
	}

	_, err = sqlS.DB()
	if err != nil {
		t.Fatal("Expect nil was %v", err)
	}

	sqlStr := "sqlite3://:memory:"
	if sqlS.String() != sqlStr {
		t.Fatalf("Expect %v was %v", sqlStr, sqlS.String())
	}
}

func Test_WrongSQLiteConfig(t *testing.T) {
	os.Clearenv()

	_, err := ReadPostgreSQL("t")
	if err == nil {
		t.Fatal("Expect error")
	}
}
