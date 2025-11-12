package logger

import "context"

type Logger interface {
	Error(ctx context.Context, err ...error)
	Info(ctx context.Context, messages ...string)
}
