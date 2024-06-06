package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SetupLogger(logLevel string) (*zap.Logger, error) {
	var level zapcore.Level
	if err := level.UnmarshalText([]byte(logLevel)); err != nil {
		return nil, err
	}

	//config := zap.Config{
	//	Level:       zap.NewAtomicLevelAt(level),
	//	Development: true,
	//	Encoding:    "json",
	//	OutputPaths: []string{"stdout"},
	//	EncoderConfig: zapcore.EncoderConfig{
	//		TimeKey:        "time",
	//		LevelKey:       "level",
	//		NameKey:        "logger",
	//		MessageKey:     "msg",
	//		CallerKey:      "caller",
	//		StacktraceKey:  "stacktrace",
	//		LineEnding:     zapcore.DefaultLineEnding,
	//		EncodeLevel:    zapcore.LowercaseLevelEncoder,
	//		EncodeTime:     zapcore.ISO8601TimeEncoder,
	//		EncodeDuration: zapcore.StringDurationEncoder,
	//		EncodeCaller:   zapcore.ShortCallerEncoder,
	//	},
	//}
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "log",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}

	atom := zap.NewAtomicLevelAt(level)
	config := zap.Config{
		Level:         atom,               // 日志级别
		Development:   true,               // 开发模式，堆栈跟踪
		Encoding:      "console",          // 输出格式 console 或 json
		EncoderConfig: encoderConfig,      // 编码器配置
		OutputPaths:   []string{"stdout"}, // 输出到指定文件 stdout（标准输出，正常颜色） stderr（错误输出，红色）
	}

	return config.Build()
}
