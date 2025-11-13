package loggerInterface

import (
	logModels "remez_story/infrastructure/logger/models"
)

type LogPublisher interface {
	SendMsg(data *logModels.LogData)
}
