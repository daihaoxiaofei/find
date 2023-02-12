package glog

import (
	"find/pkg/runvalue"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"path"
	"strings"
)

var okLog *zap.Logger

// 记录成功匹配的数据
func init() {
	logFilePath := path.Join(runvalue.RootPath, `log`, `ok.log`)
	logFilePath = logFilePath[strings.Index(logFilePath, `:`)+1:]
	cfg := zap.Config{
		Encoding:    `console`,
		Level:       zap.NewAtomicLevelAt(zapcore.DebugLevel), // 日志等级,
		OutputPaths: []string{"stderr", logFilePath},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.TimeEncoderOfLayout("01-02 15:04:05"),

			CallerKey:      "linenum",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   func(zapcore.EntryCaller, zapcore.PrimitiveArrayEncoder) {}, // 不显示调用路径
			EncodeName:     zapcore.FullNameEncoder,
			// StacktraceKey:  "stacktrace",
		},
	}
	logger, err := cfg.Build(zap.AddCallerSkip(1))
	if err != nil {
		fmt.Println("log build fail:" + err.Error())
	}
	okLog = logger
}

func Ok(msg string, fields ...zap.Field) {
	okLog.Info(msg, fields...)
}
