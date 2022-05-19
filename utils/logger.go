package utils

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//LogD will be used to log all
var LogD *zap.Logger

//StdLog will log all standard log calls to the log package
var StdLog *log.Logger

//InitializeLogger function returns a logger
func InitializeLogger() {

	cfg := zap.Config{
		Level:            zap.NewAtomicLevel(),
		Encoding:         "json",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			LineEnding:     "\n",
			LevelKey:       "level",
			MessageKey:     "message",
			NameKey:        "logger",
			TimeKey:        "time",
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeName:     zapcore.FullNameEncoder,
			EncodeDuration: zapcore.NanosDurationEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
		},
	}
	logger, _ := cfg.Build()
	logger.Sync()

	StdLog = zap.NewStdLog(logger)

	//Redirecting all standard log package logs to zap logger
	zap.RedirectStdLog(logger)

	LogD = logger
}
