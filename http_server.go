package speci

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type (
	HTTPServerSpecs struct {
		Host string `envconfig:"http_host"`
		Port int    `envconfig:"http_port"`
	}
)

func ReadHTTPServer(prefix string) (*HTTPServerSpecs, error) {
	specs := &HTTPServerSpecs{}
	err := envconfig.Process(prefix, specs)
	if err != nil {
		return nil, err
	}

	if specs.Host == "" {
		return nil, missingFieldError{"host", "HTTP-Server"}
	} else if specs.Port == 0 {
		return nil, missingFieldError{"port", "HTTP-Server"}
	}

	return specs, nil
}

func (s HTTPServerSpecs) String() string {
	return fmt.Sprintf("%v:%v", s.Host, s.Port)
}
