package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

// InitLogger 初始化日志记录器，设置日志级别和格式
func InitLogger(level string) {
	var logLevel zapcore.Level

	// 尝试设置日志级别，如果出错则默认为 Info 级别
	if err := logLevel.Set(level); err != nil {
		logLevel = zap.InfoLevel
	}

	// 配置日志记录器
	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(logLevel), // 设置日志级别
		Development: false,                          // 是否开发模式（不显示堆栈信息）
		Encoding:    "json",                         // 或者使用 console，如果你希望更易读的文本格式
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "msg",                         // 日志消息字段的键名
			LevelKey:     "level",                       // 日志级别字段的键名
			TimeKey:      "ts",                          // 时间字段的键名
			EncodeLevel:  zapcore.LowercaseLevelEncoder, // 根据需要选择是否使用颜色
			EncodeTime:   zapcore.ISO8601TimeEncoder,    // 将时间格式化为 ISO 8601 格式
			EncodeCaller: zapcore.ShortCallerEncoder,    // 确保这一行被包含，以输出调用者信息
		},
		OutputPaths:      []string{"stdout"}, // 输出路径（标准输出）
		ErrorOutputPaths: []string{"stderr"}, // 错误输出路径（标准错误输出）
	}

	// 确保使用zap.AddCaller()添加调用者信息
	var err error
	log, err = config.Build(zap.AddCaller()) // 添加调用者信息
	if err != nil {
		panic(err)
	}

	zap.ReplaceGlobals(log) // 替换全局日志实例，以便可以通过zap.L()访问
}

// GetLogger 获取日志记录器实例
func GetLogger() *zap.Logger {
	return log
}

// 便捷访问方法
func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

// Error 记录错误级别日志
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
