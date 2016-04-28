package speci

import "fmt"

type missingFieldError struct {
	field  string
	config string
}

func (err missingFieldError) Error() string {
	return fmt.Sprintf("Missing %v in %v envconfig", err.field, err.config)
}
