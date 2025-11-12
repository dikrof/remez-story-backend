package errors

import "fmt"

const UnknownErrorCode = ErrorCode("unknown_error")

type Error struct {
	code  ErrorCode
	msg   string
	level ErrorLevel
}

func NewError(code ErrorCode, msg string) *Error {
	return &Error{
		code:  code,
		msg:   msg,
		level: Levels.Error(),
	}
}

func NewErrorWithLevel(code ErrorCode, msg string, level ErrorLevel) *Error {
	return &Error{
		code:  code,
		msg:   msg,
		level: level,
	}
}

func NewErrorFrom(err error) *Error {
	return CastOrWrap(err, UnknownErrorCode)
}

func (e *Error) Code() ErrorCode {
	return e.code
}

func (e *Error) Message() string {
	return e.msg
}

func (e *Error) Level() ErrorLevel {
	return e.level
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s. Level - '%s'", e.code, e.msg, e.level)
}

func (e *Error) Equals(err *Error) bool {
	if err == nil {
		return false
	}

	return e.code == err.code && e.msg == err.msg
}
