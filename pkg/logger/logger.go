package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

var (
	log *zap.SugaredLogger
)

func Init(cfg Config) error {
	var level zapcore.Level
	if err := level.UnmarshalText([]byte(cfg.Level)); err != nil {
		return fmt.Errorf("invalid log level: %w", err)
	}

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	if cfg.DevMode {
		encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	var encoder zapcore.Encoder
	if strings.ToLower(cfg.Encoding) == "json" {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	}

	writerSyncer := zapcore.AddSync(os.Stdout)
	if cfg.OutputPath != "" && cfg.OutputPath != "stdout" {
		file, err := os.OpenFile(cfg.OutputPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return fmt.Errorf("failed to open log file: %w", err)
		}
		writerSyncer = zapcore.AddSync(file)
	}

	core := zapcore.NewCore(encoder, writerSyncer, level)

	var zapLogger *zap.Logger
	if cfg.DevMode {
		zapLogger = zap.New(core, zap.Development(), zap.AddCaller())
	} else {
		zapLogger = zap.New(core, zap.AddCaller())
	}

	log = zapLogger.Sugar()
	return nil
}

func Sugar() *zap.SugaredLogger {
	if log == nil {
		panic("logger not initialized. Call logger.Init(config) first.")
	}
	return log
}
