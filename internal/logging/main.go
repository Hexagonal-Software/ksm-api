package logging

import "github.com/gofiber/fiber/v2/log"

type Logger struct {
	logger log.Logger
}

var Log *Logger

func NewLogger(l int) {
	lgr := log.DefaultLogger()
	lgr.SetLevel(log.Level(l))

	Log = &Logger{
		logger: lgr,
	}
}

func (l *Logger) Trace(v ...interface{}) {
	l.logger.Trace(v)
}

func (l *Logger) Debug(v ...interface{}) {
	l.logger.Debug(v)
}

func (l *Logger) Info(v ...interface{}) {
	l.logger.Info(v)
}

func (l *Logger) Warn(v ...interface{}) {
	l.logger.Warn(v)
}

func (l *Logger) Error(v ...interface{}) {
	l.logger.Error(v)
}

func (l *Logger) Fatal(v ...interface{}) {
	l.logger.Fatal(v)
}
