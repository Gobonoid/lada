package lada

import "fmt"

type Error struct {
	message    string
	originator string
	error
}

func NewError(message string) Error {
	return Error {
		message:    message,
		originator: message,
	}
}

func (e Error) Originator() string {
	return e.originator
}

func (e Error) Error() string {
	if e.error != nil {
		return fmt.Sprintf("%s : %s", e.message, e.error.Error())
	}
	return e.message
}

func (e Error) CausedBy(err error) error {
	e.error = err
	return e
}

func (e Error) New(a ...interface{}) Error {
	e.message = fmt.Sprintf(e.message, a...)
	return e
}

func (e Error) Unwrap() error {
	return e.error
}

func (e Error) Cause() error {
	return e.error
}

func (e Error) Is(other error) bool {
	otherError, ok := other.(interface{ Originator() string })
	if ok {
		return e.Originator() == otherError.Originator()
	}
	return false
}

func (e Error) Message() string {
	return e.originator
}

var (
	IoReaderError                   = NewError("could not read from the source")
	IoWriterError                   = NewError("could not write to the source")
	CursorOperationError            = NewError("could not operate on cursor")
	UnexpectedWildcardArgumentError = NewError("unexpected wildcard argument `%s` in the raw `%s`")
	CommandFormatParseError     	= NewError("failed to parse pattern `%s`")
	UnexpectedArgumentError         = NewError("unexpected argument `%s` in command `%s`")
	UnknownArgumentError			= NewError("unknown argument `%s`")
	MissingArgumentValueError		= NewError("argument `%s` expects a value to be passed")
	UnexpectedArgumentValue			= NewError("argument `%s` expects no value")
	InvalidArgumentValueError		= NewError("argument `%s` has invalid value `%s`")
	ArgumentAlreadyDefinedError		= NewError("argument `%s` is already defined in the raw `%s`")
	InvalidArgumentNameError		= NewError("invalid argument's name `%s`")
	CommandError					= NewError("there was an error while executing command `%s`")
)
