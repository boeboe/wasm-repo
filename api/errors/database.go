package errors

import "fmt"

type DatabaseError struct {
	Operation string
	Err       error
}

func (e *DatabaseError) Error() string {
	return fmt.Sprintf("Database error during %s: %v", e.Operation, e.Err)
}
