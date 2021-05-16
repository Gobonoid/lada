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
	return fmt.Sprintf("%v : %v", e.Message(), e.error.Error())
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
	InvalidCommandIdentifierError   = NewError("invalid identifier name `%s` in definition")
	UnexpectedCommandParameterError = NewError("unexpected parameter `%s` in definition `%s`")
	UnexpectedWildcardArgumentError = NewError("unexpected wildcard argument `%s` in definition `%s`")
	CommandDefinitionParseError     = NewError("failed to parse definition `%s`")
	CommandNameDoesNotMatchError    = NewError("command name `%s` does not match `%s`")
	UnexpectedArgumentError         = NewError("unexpected argument `%s` in command `%s`")
	UnexpectedParameterError        = NewError("unexpected parameter `%s` in command `%s`")
	UnexpectedFlagValueError		= NewError("flags `%s` accepts no value")
	UnknownParameterError			= NewError("unknown parameter `%s` in comamand `%s`")
	MissingParameterValueError		= NewError("parameter `%s` expects a value to be passed")
	CommandError					= NewError("there was an error while executing command `%s`")
)
