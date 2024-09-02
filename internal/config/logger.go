package config

import (
	"fmt"
	"go.uber.org/zap"
	"log"
)

type Logger struct {
	zap *zap.Logger
}

func NewLogger() *Logger {
	l, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}
	return &Logger{zap: l}
}

func (l *Logger) Debug(format string, a ...any) {
	l.zap.Debug(fmt.Sprintf(format, a...))
}

func (l *Logger) Info(format string, a ...any) {
	l.zap.Info(fmt.Sprintf(format, a...))
}

func (l *Logger) Warn(format string, a ...any) {
	l.zap.Warn(fmt.Sprintf(format, a...))
}

func (l *Logger) Error(format string, a ...any) {
	l.zap.Error(fmt.Sprintf(format, a...))
}
