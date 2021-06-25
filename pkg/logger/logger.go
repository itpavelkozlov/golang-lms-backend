package logger

import (
	"github.com/itpavelkozlov/golang-lms-backend/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type Logger interface {
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
}

var levelMap = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"panic": zapcore.PanicLevel,
	"fatal": zapcore.FatalLevel,
}

func getLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}

type LoggerImp struct {
	logger *zap.Logger
}

func (l LoggerImp) Debug(msg string, fields ...zap.Field) {
	l.logger.Debug(msg, fields...)
}

func (l LoggerImp) Info(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}

func (l LoggerImp) Warn(msg string, fields ...zap.Field) {
	l.logger.Warn(msg, fields...)
}

func (l LoggerImp) Error(msg string, fields ...zap.Field) {
	l.logger.Error(msg, fields...)
}

func NewLogger(config *config.Config) (Logger, error) {

	level := getLoggerLevel(config.Service.Logger.Level)
	syncWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   config.Service.Logger.Filename,
		MaxSize:    config.Service.Logger.MaxSize,
		MaxBackups: config.Service.Logger.MaxBackups,
		LocalTime:  true,
		Compress:   config.Service.Logger.Compress,
	})
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder

	cores := make([]zapcore.Core, 0)
	for i := range config.Service.Logger.Cores {
		switch config.Service.Logger.Cores[i] {
		case "file":
			cores = append(cores, zapcore.NewCore(zapcore.NewJSONEncoder(encoder), syncWriter, zap.NewAtomicLevelAt(level)))
		case "stdout":
			cores = append(cores, zapcore.NewCore(zapcore.NewJSONEncoder(encoder), zapcore.AddSync(os.Stdout), level))
		case "stderr":
			cores = append(cores, zapcore.NewCore(zapcore.NewJSONEncoder(encoder), zapcore.AddSync(os.Stderr), level))
		}
	}

	core := zapcore.NewTee(cores...)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	logger.Info(
		"Logger started",
		zap.String("Filename", config.Service.Logger.Filename),
		zap.Int("Max file size", config.Service.Logger.MaxSize),
		zap.Int("Max file backups", config.Service.Logger.MaxBackups),
		zap.String("Level", config.Service.Logger.Level),
	)

	return &LoggerImp{
		logger: logger,
	}, nil
}
