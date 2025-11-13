package logModels

import (
	"context"
	"fmt"
)

type LogLevel int8

const (
	DebugLevel LogLevel = iota - 1
	InfoLevel
	WarnLevel
	ErrorLevel
	DPanicLevel
	PanicLevel
	FatalLevel
)

func (l LogLevel) String() string {
	switch l {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	case DPanicLevel:
		return "dpanic"
	case PanicLevel:
		return "panic"
	case FatalLevel:
		return "fatal"
	default:
		return fmt.Sprintf("Level(%d)", l)
	}
}

const (
	FieldErrKey       = "error"
	FieldComponentKey = "component"
	FieldFilenameKey  = "filename"
)

type LogData struct {
	Ctx    context.Context
	Msg    string
	Fields []*LogField
	Level  LogLevel
}

type LogField struct {
	Key     string
	Integer int
	Float   float64
	String  string
	Object  interface{}
}
