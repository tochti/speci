package speci

import (
	"os"
	"testing"
)

func TestReadServer(t *testing.T) {
	os.Clearenv()

	os.Setenv("HTTP_HOST", "127.0.0.1")
	os.Setenv("HTTP_PORT", "9090")

	srv, err := ReadHTTPServer()

	if err != nil {
		t.Fatal(err)
	}

	exp := "127.0.0.1:9090"
	if exp != srv.String() {
		t.Fatalf("Expect %v was %v", exp, srv)
	}

}

func TestReadServer_WrongConfig(t *testing.T) {
	os.Clearenv()

	_, err := ReadHTTPServer()

	if err != ErrHTTPServer {
		t.Fatal("Expect %v", ErrHTTPServer.Error())
	}

}
