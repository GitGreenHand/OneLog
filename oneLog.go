package oneLog

import (
	"fmt"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"strings"
)



type OneLogger struct {
	output     zerolog.ConsoleWriter
	fileLogger *lumberjack.Logger
	OutLogger  zerolog.Logger
	writer     io.Writer
	MaxSize    int // megabytes
	MaxBackups int
	MaxAge     int  //days
	Compress   bool // disabled by default
}

func (m *OneLogger) Build() *OneLogger {
	m.output.Out = m.writer
	m.OutLogger.Output(m.output)
	return m
}

func (m *OneLogger) WithConsoleWriter() *OneLogger {
	m.writer = os.Stdout
	return m
}

// WithTimeFormat /
func (m *OneLogger) WithTimeFormat(timeFormat string) *OneLogger {
	m.output.TimeFormat = timeFormat
	return m
}

// WithLevel /
func (m *OneLogger) WithLevel(level string) *OneLogger {
	m.output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf(level, i))
	}
	return m
}

// WithFieldName /
func (m *OneLogger) WithFieldName(FieldName string) *OneLogger {

	m.output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf(FieldName, i)
	}
	return m
}

// WithFieldValue /
func (m *OneLogger) WithFieldValue(FieldValue string) *OneLogger {

	m.output.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf(FieldValue, i)
	}
	return m
}

// WithMessage /
func (m *OneLogger) WithMessage(message string) *OneLogger {
	m.output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf(message, i)
	}
	return m
}

// WithColor /
func (m *OneLogger) WithColor(colorEnable bool) *OneLogger {
	m.output.NoColor = !colorEnable
	return m
}

// WithFileWriter /
func (m *OneLogger) WithFileWriter(fileName string, consoleEnable bool) *OneLogger {
	dir, _ := os.Getwd()
	m.fileLogger = &lumberjack.Logger{
		Filename:   dir + fileName,
		MaxSize:    m.MaxSize, // megabytes
		MaxBackups: m.MaxBackups,
		MaxAge:     m.MaxAge,   //days
		Compress:   m.Compress, // disabled by default
	}
	if consoleEnable {
		m.writer = io.MultiWriter(m.fileLogger, os.Stdout)
	} else {
		m.writer = io.MultiWriter(m.fileLogger)
	}
	return m
}

func NewDefaultConsoleLog() *OneLogger {
	oneLogger := new(OneLogger)
	output := oneLogger.output
	output.Out = os.Stdout
	output.TimeFormat = "2006-01-02 15:04:05"
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("[%s]", i))
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("[%s", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf("%s]", i)
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	oneLogger.OutLogger = zerolog.New(output).With().Timestamp().Logger()
	return oneLogger
}
