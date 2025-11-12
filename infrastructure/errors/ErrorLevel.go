package errors

import "fmt"

var (
	UnsupportedLevelErrorCode ErrorCode = "dcd82a3f-001"
)

func ErrUnsupportedLevel(levelStr string) error {
	errMessage := fmt.Sprintf("Unsupported ErrorLevel = '%s'", levelStr)
	return NewError(UnsupportedLevelErrorCode, errMessage)
}

type ErrorLevel string

func (l ErrorLevel) String() string {
	return string(l)
}

const (
	infoErrorLevel     = "info"
	warnErrorLevel     = "warn"
	errorErrorLevel    = "error"
	criticalErrorLevel = "critical"
)

type ErrorLevelEnum map[string]ErrorLevel

var Levels = ErrorLevelEnum{
	infoErrorLevel:     infoErrorLevel,
	warnErrorLevel:     warnErrorLevel,
	errorErrorLevel:    errorErrorLevel,
	criticalErrorLevel: criticalErrorLevel,
}

func (e ErrorLevelEnum) Info() ErrorLevel {
	return e[infoErrorLevel]
}

func (e ErrorLevelEnum) Warn() ErrorLevel {
	return e[warnErrorLevel]
}

func (e ErrorLevelEnum) Error() ErrorLevel {
	return e[errorErrorLevel]
}

func (e ErrorLevelEnum) Critical() ErrorLevel {
	return e[criticalErrorLevel]
}

func (e ErrorLevelEnum) Of(levelCode string) (ErrorLevel, error) {
	value, ok := e[levelCode]
	if !ok {
		return "", ErrUnsupportedLevel(levelCode)
	}

	return value, nil
}
