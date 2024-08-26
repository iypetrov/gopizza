package logger

import (
	"go.uber.org/zap"
	"log"
)

type Logger interface {
	Debug(event interface{})
	Info(event interface{})
	Warn(event interface{})
	Error(event interface{})
}

type UberZapLogger struct {
	zap *zap.Logger
}

func New() *Logger {
	l, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}
	var logger Logger = &UberZapLogger{zap: l}
	return &logger
}

func (l *UberZapLogger) Debug(event interface{}) {
	l.zap.Debug("received debug event", zap.Any("event", event))
}

func (l *UberZapLogger) Info(event interface{}) {
	l.zap.Info("received info event", zap.Any("event", event))
}

func (l *UberZapLogger) Warn(event interface{}) {
	l.zap.Warn("received warn event", zap.Any("event", event))
}

func (l *UberZapLogger) Error(event interface{}) {
	l.zap.Error("received error event", zap.Any("event", event))
}
