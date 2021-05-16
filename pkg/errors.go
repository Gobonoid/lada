package lada

import "fmt"

type Error struct {
	message string
	error
}

func NewError(message string) Error {
	return Error{
		message: message,
	}
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
	otherError, ok := other.(interface{ Message() string })
	if ok {
		return e.Message() == otherError.Message()
	}
	return false
}

func (e Error) Message() string {
	return e.message
}

var (
	IoReaderError                   = NewError("could not read from the source")
	IoWriterError                   = NewError("could not write to the source")
	CursorOperationError            = NewError("could not operate on cursor")
	InvalidCommandIdentifierError   = NewError("invalid identifier name `%s` in definition")
	UnexpectedCommandParameterError = NewError("unexpected parameter `%s` in definition `%s`")
	UnexpectedWildcardArgumentError = NewError("unexpected wildcard argument `%s` in definition `%s`")
	CommandDefinitionParseError     = NewError("failed to parse definition `%s`")
	CannotCreateCommandError 		= NewError("cannot create command")
	CommandNameDoesNotMatchError	= NewError("command name `%s` does not match `%s`")
	UnexpectedArgument				= NewError("unexpected argument `%s` in command `%s`")
)
