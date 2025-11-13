package loggerInterface

import (
	"context"
	logModels "remez_story/infrastructure/logger/models"
)

type Logger interface {
	Error(ctx context.Context, err error, options ...logModels.Option)
	Errors(ctx context.Context, errs []error, options ...logModels.Option)
	Info(ctx context.Context, message string, options ...logModels.Option)
	Warning(ctx context.Context, message string, options ...logModels.Option)
	Debug(ctx context.Context, message string, options ...logModels.Option)
}
