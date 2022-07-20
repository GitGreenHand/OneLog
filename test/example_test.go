package test

import (
	"OneLog"
	"testing"
)

func TestOneLog(t *testing.T) {
	logger1 := oneLog.NewDefaultConsoleLog()
	logger1.OutLogger.Info().Msg("default logs")

	// 自定logger
	logger2 := new(oneLog.OneLogger).
		WithTimeFormat("2006-01-02 15:04:05").
		WithLevel("[logger2] %s [record]").
		WithMessage("{ %s }").
		WithColor(false).
		WithConsoleWriter().
		Build()
	logger2.OutLogger.Info().Msg("logger2 ..........")

	logger3 := new(oneLog.OneLogger).
		WithTimeFormat("2006-01-02 15:04:05").
		WithLevel("[logger3] %s [record]").
		WithMessage("{ %s }").
		WithColor(false).
		WithFileWriter("/test/test.log", true).
		Build()

	logger3.OutLogger.Info().Msg("logger3.......")

	logger4 := new(oneLog.OneLogger).
		WithTimeFormat("2006-01-02 15:04:05").
		WithLevel("[logger3] %s [record]").
		WithMessage("{ %s }").
		WithColor(false).
		WithFileWriter("/test.log", true)
	logger4.MaxSize = 200
	logger4.MaxAge = 20
	logger4.Build()
	logger4.OutLogger.Info().Msg("logger4 ..........")
}
