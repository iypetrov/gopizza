package logger

import (
	"fmt"
	"go.uber.org/zap"
	"log"
)

type Logger interface {
	Debug(format string, a ...any)
	Info(format string, a ...any)
	Warn(format string, a ...any)
	Error(format string, a ...any)
}

type UberZapLogger struct {
	zap *zap.Logger
}

func New() Logger {
	l, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}
	return &UberZapLogger{zap: l}
}

func (l *UberZapLogger) Debug(format string, a ...any) {
	l.zap.Debug("received debug event", zap.Any("event", fmt.Sprintf(format, a...)))
}

func (l *UberZapLogger) Info(format string, a ...any) {
	l.zap.Info("received info event", zap.Any("event", fmt.Sprintf(format, a...)))
}

func (l *UberZapLogger) Warn(format string, a ...any) {
	l.zap.Warn("received warn event", zap.Any("event", fmt.Sprintf(format, a...)))
}

func (l *UberZapLogger) Error(format string, a ...any) {
	l.zap.Error("received error event", zap.Any("event", fmt.Sprintf(format, a...)))
}
