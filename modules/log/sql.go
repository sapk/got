package log

import (
	"log"

	"github.com/rs/zerolog"
	"xorm.io/core"
)

//NewSQLLogger create a new logger for xorm
func NewSQLLogger(l *zerolog.Logger) *SQLLogger {
	return &SQLLogger{l}
}

type SQLLogger struct {
	logger *zerolog.Logger
}

func (l *SQLLogger) Debug(v ...interface{}) {
	l.logger.Print(v...) //TODO debug level
}

func (l *SQLLogger) Debugf(format string, v ...interface{}) {
	l.logger.Printf(format, v...) //TODO debug level
}

func (l *SQLLogger) Error(v ...interface{}) {
	l.logger.Print(v...) //TODO Error level
}

func (l *SQLLogger) Errorf(format string, v ...interface{}) {
	l.logger.Printf(format, v...) //TODO Error level
}

func (l *SQLLogger) Info(v ...interface{}) {
	l.logger.Print(v...) //TODO Info level
}

func (l *SQLLogger) Infof(format string, v ...interface{}) {
	l.logger.Printf(format, v...) //TODO Info level
}

func (l *SQLLogger) Warn(v ...interface{}) {
	l.logger.Print(v...) //TODO Warn level
}

func (l *SQLLogger) Warnf(format string, v ...interface{}) {
	log.Printf(format, v...) //TODO Warn level
}

func (l *SQLLogger) Level() core.LogLevel {
	return core.LOG_DEBUG //TODO
}

func (l *SQLLogger) SetLevel(lvl core.LogLevel) {
	l.logger.Debug().Msgf("xorm.log.SetLevel %d", lvl)
	//zerolog.SetGlobalLevel(zerolog.DebugLevel)
}
func (l *SQLLogger) ShowSQL(show ...bool) {
	//TODO
}

func (l *SQLLogger) IsShowSQL() bool {
	return true
}

//*/
