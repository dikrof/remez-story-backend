package zapLogger

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	logModels "remez_story/infrastructure/logger/models"
)

const (
	timeTag = "timestamp"
)

type ZapLogger struct {
	*zap.Logger
	appID string
	env   string
}

func NewZapLogger(appID, env string) *ZapLogger {
	config := getEncoderConfig()
	coreConsole := zapcore.NewCore(zapcore.NewJSONEncoder(config), os.Stdout, getAllLevelFunc())
	zapLogger := zap.New(zapcore.NewTee(coreConsole))

	return &ZapLogger{
		Logger: zapLogger,
		appID:  appID,
		env:    env,
	}
}

func (l *ZapLogger) SendMsg(logData *logModels.LogData) {
	if logData.Ctx == nil {
		logData.Ctx = context.Background()
	}

	fields := []zapcore.Field{
		zap.String("service_name", l.appID),
		zap.String("env", l.env),
	}

	resFields := l.getPayloadFields(logData)
	fields = append(fields, resFields...)

	switch logData.Level {
	case logModels.ErrorLevel:
		l.Error(logData.Msg, fields...)
	case logModels.WarnLevel:
		l.Warn(logData.Msg, fields...)
	case logModels.InfoLevel:
		l.Info(logData.Msg, fields...)
	case logModels.DebugLevel:
		l.Debug(logData.Msg, fields...)
	case logModels.FatalLevel:
		l.Fatal(logData.Msg, fields...)
	}
}

func (l *ZapLogger) getPayloadFields(logData *logModels.LogData) []zap.Field {
	var resFields []zap.Field
	resFields = append(resFields, zap.Namespace("payload"))
	for _, f := range logData.Fields {
		if f.Integer != 0 {
			resFields = append(resFields, zap.Int(f.Key, f.Integer))
		}
		if f.String != "" {
			resFields = append(resFields, zap.String(f.Key, f.String))
		}
		if f.Float != 0.0 {
			resFields = append(resFields, zap.Float64(f.Key, f.Float))
		}
		if f.Object != nil {
			resFields = append(resFields, zap.Any(f.Key, f.Object))
		}
	}
	return resFields
}

func getEncoderConfig() zapcore.EncoderConfig {
	config := zap.NewProductionEncoderConfig()
	config.TimeKey = timeTag
	config.EncodeTime = zapcore.RFC3339TimeEncoder
	return config
}

func getAllLevelFunc() zap.LevelEnablerFunc {
	return func(l zapcore.Level) bool { return true }
}
