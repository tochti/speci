package speci

import (
	"errors"
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

var (
	ErrHTTPServer = errors.New("Wrong HTTP-Server config")
)

type (
	HTTPServerSpecs struct {
		Host string `envconfig:"http_host"`
		Port int    `envconfig:"http_port"`
	}
)

func ReadHTTPServer() (*HTTPServerSpecs, error) {
	specs := &HTTPServerSpecs{}
	err := envconfig.Process(AppName, specs)
	if err != nil {
		return nil, err
	}

	if specs.Host == "" {
		return nil, ErrHTTPServer
	}

	return specs, nil
}

func (s HTTPServerSpecs) String() string {
	return fmt.Sprintf("%v:%v", s.Host, s.Port)
}
