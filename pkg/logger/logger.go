package logger

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Config defines options for logger creation
// type Config struct {
// 	Filename   string
// 	MaxSize    int
// 	MaxBackups int
// 	MaxAge     int
// 	Compress   bool
// 	LogLevel   zapcore.LevelEnabler
// }

// NewLogger creates a new zap logger with the provided configuration
// func NewLogger(cfg Config) (*zap.Logger, error) {
func NewLogger(LogFilename string, MaxSize int, MaxBackups int, MaxAge int, Compress bool, LogLevel zapcore.LevelEnabler) (*zap.Logger, error) {
	var writer zapcore.WriteSyncer

	// Use lumberjack for log rotation
	writer = zapcore.AddSync(io.MultiWriter(&lumberjack.Logger{
		Filename:   LogFilename,
		MaxSize:    MaxSize,
		MaxBackups: MaxBackups,
		MaxAge:     MaxAge,
		Compress:   Compress,
	}, os.Stdout))
	// Default to stdout

	encoderCfg := zap.NewDevelopmentEncoderConfig() // Adjust encoder as needed
	encoder := zapcore.NewJSONEncoder(encoderCfg)

	core := zapcore.NewCore(
		encoder,
		writer,
		LogLevel, // Set log level if provided
	)

	return zap.New(core), nil
}
