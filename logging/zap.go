package logging

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"papertrader.io/backoffice/config"
)

func InitLogger(configuration *config.GeneralConfig) *zap.Logger {
	writerSyncer := getLogWriter(configuration)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	return logger
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(configuration *config.GeneralConfig) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename: configuration.Logging.Default.File,
		MaxSize: 10,
		MaxBackups: 5,
		MaxAge: 30,
		Compress: false,
	}
	return zapcore.AddSync(lumberJackLogger)
}
