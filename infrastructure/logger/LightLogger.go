package logger

import (
	"context"
	"log"
)

type LightLogger struct{}

func NewLightLogger() *LightLogger {
	return &LightLogger{}
}

func (l *LightLogger) Error(ctx context.Context, errs ...error) {
	for _, err := range errs {
		log.Printf("ERROR: %v\n", err)
	}
}

func (l *LightLogger) Info(ctx context.Context, messages ...string) {
	for _, msg := range messages {
		log.Printf("INFO: %s\n", msg)
	}
}
