package errors

import "fmt"

type FileError struct {
	File      string
	Operation string
	Err       error
}

func (e *FileError) Error() string {
	return fmt.Sprintf("File error during %s on %s: %v", e.Operation, e.File, e.Err)
}
