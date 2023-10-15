package errors

import "fmt"

type JSONDecodingError struct {
	Source string
	Err    error
}

func (e *JSONDecodingError) Error() string {
	return fmt.Sprintf("JSON decoding error in %s: %v", e.Source, e.Err)
}
