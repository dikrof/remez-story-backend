package logger

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	loggerInterface "remez_story/infrastructure/logger/interface"
	logModels "remez_story/infrastructure/logger/models"
	"strings"
)

type Logger struct {
	logChan chan<- *logModels.LogData
}

func NewLogger(logChan chan<- *logModels.LogData) *Logger {
	return &Logger{logChan: logChan}
}

func (l *Logger) Error(ctx context.Context, err error, options ...logModels.Option) {
	opts := &logModels.Options{}
	for _, opt := range options {
		opt(opts)
	}
	l.error(ctx, err, opts)
}

func (l *Logger) Errors(ctx context.Context, errs []error, options ...logModels.Option) {
	opts := &logModels.Options{}
	for _, opt := range options {
		opt(opts)
	}
	for _, err := range errs {
		l.error(ctx, err, opts)
	}
}

func (l *Logger) error(ctx context.Context, err error, opts *logModels.Options) {
	extendedErr := errors.WithStack(err)
	logData := &logModels.LogData{
		Ctx:    ctx,
		Msg:    extendedErr.Error(),
		Fields: []*logModels.LogField{},
		Level:  logModels.ErrorLevel,
	}

	if opts.WithStackTrace() {
		var fileNames []string
		if stackTracerErr, ok := extendedErr.(loggerInterface.StackTracer); ok {
			stacktrace := stackTracerErr.StackTrace()
			if len(stacktrace) > 0 {
				for i := 1; i < len(stacktrace); i++ {
					fileNames = append(fileNames, fmt.Sprintf("%s:%d", stacktrace[i], stacktrace[i]))
				}
			}
		}
		logData.Fields = append(logData.Fields,
			&logModels.LogField{Key: logModels.FieldFilenameKey, String: strings.Join(fileNames, " <- ")})
	}

	if len(opts.GetFields()) > 0 {
		logData.Fields = append(logData.Fields, opts.GetFields()...)
	}
	if opts.GetComponent() != "" {
		logData.Fields = append(logData.Fields,
			&logModels.LogField{Key: logModels.FieldComponentKey, String: opts.GetComponent()})
	}

	go l.sendData(logData)
}

func (l *Logger) Warning(ctx context.Context, message string, options ...logModels.Option) {
	l.logMsg(ctx, logModels.WarnLevel, message, options...)
}

func (l *Logger) Info(ctx context.Context, message string, options ...logModels.Option) {
	l.logMsg(ctx, logModels.InfoLevel, message, options...)
}

func (l *Logger) Debug(ctx context.Context, message string, options ...logModels.Option) {
	l.logMsg(ctx, logModels.DebugLevel, message, options...)
}

func (l *Logger) logMsg(ctx context.Context, level logModels.LogLevel, message string, options ...logModels.Option) {
	opts := &logModels.Options{}
	for _, opt := range options {
		opt(opts)
	}

	logData := &logModels.LogData{
		Ctx:    ctx,
		Msg:    message,
		Fields: opts.GetFields(),
		Level:  level,
	}

	if opts.GetComponent() != "" {
		logData.Fields = append(logData.Fields,
			&logModels.LogField{Key: logModels.FieldComponentKey, String: opts.GetComponent()})
	}

	go l.sendData(logData)
}

func (l *Logger) sendData(logData *logModels.LogData) {
	l.logChan <- logData
}
