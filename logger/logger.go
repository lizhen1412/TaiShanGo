package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func InitLogger(level string) {
	var logLevel zapcore.Level
	if err := logLevel.Set(level); err != nil {
		logLevel = zap.InfoLevel
	}

	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(logLevel),
		Development: false,
		Encoding:    "json", // 或者使用 console，如果你希望更易读的文本格式
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "msg",
			LevelKey:     "level",
			TimeKey:      "ts",
			EncodeLevel:  zapcore.LowercaseLevelEncoder, // 根据需要选择是否使用颜色
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder, // 确保这一行被包含，以输出调用者信息
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	// 确保使用zap.AddCaller()添加调用者信息
	var err error
	log, err = config.Build(zap.AddCaller())
	if err != nil {
		panic(err)
	}

	zap.ReplaceGlobals(log) // 替换全局日志实例，以便可以通过zap.L()访问
}

func GetLogger() *zap.Logger {
	return log
}

// 便捷访问方法
func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}

// 添加Debug级别日志方法
func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

// 添加Warn级别日志方法
func Warn(message string, fields ...zap.Field) {
	log.Warn(message, fields...)
}

// 添加Fatal级别日志方法
func Fatal(message string, fields ...zap.Field) {
	log.Fatal(message, fields...)
}

// 添加Panic级别日志方法，会触发panic
func Panic(message string, fields ...zap.Field) {
	log.Panic(message, fields...)
}
