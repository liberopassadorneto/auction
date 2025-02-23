package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.Logger
)

func init() {
	logConfiguration := zap.Config{
		Level:    zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			LevelKey:     "level",
			TimeKey:      "time",
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	log, _ = logConfiguration.Build()
}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
	log.Sync()
}

func Error(message string, err error, fields ...zap.Field) {
	fieldsWithError := append(fields, zap.NamedError("error", err))
	log.Error(message, fieldsWithError...)
	log.Sync()
}
