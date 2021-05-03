package lada

import "fmt"

type ladaError struct {
	message string
	error
}

func (e ladaError) Error() string {
	return fmt.Sprintf("%v : %v", e.message, e.error.Error())
}

func (e ladaError) causedBy(err error) error {
	e.error = err
	return e
}

func (e ladaError) Unwrap() error {
	return e.error
}

func (e ladaError) Cause() error {
	return e.error
}

func (e ladaError) Is(other error) bool {
	otherError, ok := other.(interface { Message() string})
	if ok {
		return e.Message() == otherError.Message()
	}
	return false
}

func (e ladaError) Message() string {
	return e.message
}

var (
	IoReaderError         = ladaError{"could not read from the source", nil}
	IoWriterError         = ladaError{"could not write to the source", nil}
	CursorOperationError  = ladaError{"could not operate on cursor", nil}
	CursorOutOfReachError = ladaError{"out of reach", nil}
)