package logger

import "go.uber.org/zap"

type Logger struct {
	logger *zap.Logger
}

func NewLogger() *Logger {
	logger, _ := zap.NewProduction()
	return &Logger{logger: logger}
}

func (l *Logger) Info(msg string, args ...interface{}) {
	l.logger.Sugar().Infof(msg, args...)
}

func (l *Logger) Warn(msg string, err error) {
	l.logger.Sugar().Warnw(msg, "error", err)
}

func (l *Logger) Fatal(msg string, err error) {
	l.logger.Sugar().Fatalw(msg, "error", err)
}

func (l *Logger) Sync() error {
	return l.logger.Sync()
}
