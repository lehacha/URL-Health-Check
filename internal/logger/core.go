package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger() *zap.Logger {

	encoderConfig := zap.NewProductionConfig()
	encoderConfig.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	// timestamp of conversion encoding

	logger, err := encoderConfig.Build()
	if err != nil {
		errorMessage := fmt.Sprintf("Error creating logger: %s", err.Error())
		panic(errorMessage)
	}

	defer logger.Sync()

	zap.ReplaceGlobals(logger)

	return logger
}
