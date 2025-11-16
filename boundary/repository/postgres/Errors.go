package postgresRepository

import (
	"context"
	"fmt"
	"remez_story/infrastructure/errors"
	loggerInterface "remez_story/infrastructure/logger/interface"
)

var (
	ErrNodeNotFound      = errors.NewErrorWithLevel("feab43fa-001", "Node not found", errors.Levels.Info())
	ErrStartNodeNotFound = errors.NewErrorWithLevel("feab43fa-002", "Start node for chapter not found", errors.Levels.Info())
	ErrFindNodes         = errors.NewErrorWithLevel("feab43fa-003", "Failed to find nodes", errors.Levels.Error())
)

var (
	ErrChapterNotFound = errors.NewErrorWithLevel("a1b2c3d4-001", "Chapter not found", errors.Levels.Info())
	ErrFindChapters    = errors.NewErrorWithLevel("a1b2c3d4-002", "Failed to find chapters", errors.Levels.Error())
)

var (
	ErrProgressNotFound = errors.NewErrorWithLevel("b2c3d4e5-001", "Progress not found", errors.Levels.Info())
	ErrSaveProgress     = errors.NewErrorWithLevel("b2c3d4e5-002", "Failed to save progress", errors.Levels.Error())
	ErrUpdateProgress   = errors.NewErrorWithLevel("b2c3d4e5-003", "Failed to update progress", errors.Levels.Error())
)

var (
	ErrEventNotFound = errors.NewErrorWithLevel("c3d4e5f6-001", "Event not found", errors.Levels.Info())
	ErrFindEvents    = errors.NewErrorWithLevel("c3d4e5f6-002", "Failed to find events", errors.Levels.Error())
)

type errorProcessor struct {
	logger loggerInterface.Logger
}

func newErrorProcessor(logger loggerInterface.Logger) *errorProcessor {
	return &errorProcessor{logger: logger}
}

func (p *errorProcessor) LogAndReturnErrNodeNotFound(ctx context.Context, nodeID int64) error {
	detailErrMsg := fmt.Sprintf("Node by id = %d not found", nodeID)
	detailError := errors.NewError(ErrNodeNotFound.Code(), detailErrMsg)
	p.logger.Error(ctx, detailError)
	return ErrNodeNotFound
}

func (p *errorProcessor) LogAndReturnErrStartNodeNotFound(ctx context.Context, chapterID int64) error {
	detailErrMsg := fmt.Sprintf("Start node for chapter ID=%d not found", chapterID)
	detailError := errors.NewError(ErrStartNodeNotFound.Code(), detailErrMsg)
	p.logger.Error(ctx, detailError)
	return ErrStartNodeNotFound
}

func (p *errorProcessor) LogAndReturnErrFindNodes(ctx context.Context, cause error) error {
	detailErrMsg := fmt.Sprintf("Failed to find nodes. Cause: %q", cause.Error())
	detailError := errors.NewError(ErrFindNodes.Code(), detailErrMsg)
	p.logger.Error(ctx, detailError)
	return ErrFindNodes
}

func (p *errorProcessor) LogAndReturnErrChapterNotFound(ctx context.Context, chapterID int64) error {
	detailErrMsg := fmt.Sprintf("Chapter by id = %d not found", chapterID)
	detailError := errors.NewError(ErrChapterNotFound.Code(), detailErrMsg)
	p.logger.Error(ctx, detailError)
	return ErrChapterNotFound
}

func (p *errorProcessor) LogAndReturnErrFindChapters(ctx context.Context, cause error) error {
	detailErrMsg := fmt.Sprintf("Failed to find chapters. Cause: %q", cause.Error())
	detailError := errors.NewError(ErrFindChapters.Code(), detailErrMsg)
	p.logger.Error(ctx, detailError)
	return ErrFindChapters
}

func (p *errorProcessor) LogAndReturnErrProgressNotFound(ctx context.Context, playerID string) error {
	detailErrMsg := fmt.Sprintf("Progress for player = %q not found", playerID)
	detailError := errors.NewError(ErrProgressNotFound.Code(), detailErrMsg)
	p.logger.Error(ctx, detailError)
	return ErrProgressNotFound
}

func (p *errorProcessor) LogAndReturnErrSaveProgress(ctx context.Context, cause error) error {
	detailErrMsg := fmt.Sprintf("Failed to save progress. Cause: %q", cause.Error())
	detailError := errors.NewError(ErrSaveProgress.Code(), detailErrMsg)
	p.logger.Error(ctx, detailError)
	return ErrSaveProgress
}

func (p *errorProcessor) LogAndReturnErrUpdateProgress(ctx context.Context, playerID string, cause error) error {
	detailErrMsg := fmt.Sprintf("Failed to update progress for player=%q. Cause: %q", playerID, cause.Error())
	detailError := errors.NewError(ErrUpdateProgress.Code(), detailErrMsg)
	p.logger.Error(ctx, detailError)
	return ErrUpdateProgress
}

func (p *errorProcessor) LogAndReturnErrEventNotFound(ctx context.Context, eventID int64) error {
	detailErrMsg := fmt.Sprintf("Event by id = %d not found", eventID)
	detailError := errors.NewError(ErrEventNotFound.Code(), detailErrMsg)
	p.logger.Error(ctx, detailError)
	return ErrEventNotFound
}

func (p *errorProcessor) LogAndReturnErrFindEvents(ctx context.Context, cause error) error {
	detailErrMsg := fmt.Sprintf("Failed to find events. Cause: %q", cause.Error())
	detailError := errors.NewError(ErrFindEvents.Code(), detailErrMsg)
	p.logger.Error(ctx, detailError)
	return ErrFindEvents
}
