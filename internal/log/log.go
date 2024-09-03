package log

import (
	"fmt"
	"go.uber.org/zap"
	"log"
)

var (
	logger *zap.Logger
)

func init() {
	l, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to create log: %v", err)
	}
	logger = l
}

func Debug(format string, a ...any) {
	logger.Debug(fmt.Sprintf(format, a...))
}

func Info(format string, a ...any) {
	logger.Info(fmt.Sprintf(format, a...))
}

func Warn(format string, a ...any) {
	logger.Warn(fmt.Sprintf(format, a...))
}

func Error(format string, a ...any) {
	logger.Error(fmt.Sprintf(format, a...))
}
