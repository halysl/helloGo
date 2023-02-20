package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"os"
)

var Logger *zap.SugaredLogger

func InitLogger() {
	// 获取 core 需要的 encoder、writeSyncer、logger
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	writeSyncer := zapcore.AddSync(os.Stdout)
	level := zap.InfoLevel
	// 提供可用 logger
	core := zapcore.NewCore(encoder, writeSyncer, level)
	_logger := zap.New(core, zap.AddCaller())
	Logger = _logger.Sugar()
}
