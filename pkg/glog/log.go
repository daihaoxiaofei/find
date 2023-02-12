package glog

import (
	"find/pkg/config"
	"find/pkg/runvalue"
	"fmt"
	"path"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	gLog *zap.Logger
)

func init() {
	var nowLevel zapcore.Level
	switch config.Run.LogLevel {
	case `info`:
		nowLevel = zapcore.InfoLevel
	case `error`:
		nowLevel = zapcore.ErrorLevel
	default:
		nowLevel = zapcore.DebugLevel
	}

	logFilePath := path.Join(runvalue.RootPath, `log`, time.Now().Format(`2006-01-02`)+`.log`)
	logFilePath = logFilePath[strings.Index(logFilePath, `:`)+1:]
	cfg := zap.Config{
		Encoding:         `console`,
		Level:            zap.NewAtomicLevelAt(nowLevel), // 日志等级,
		OutputPaths:      []string{"stderr", logFilePath},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.TimeEncoderOfLayout("01-02 15:04:05"),

			CallerKey:      "linenum",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.FullCallerEncoder, // zapcore.ShortCallerEncoder, // 短路径编码器
			EncodeName:     zapcore.FullNameEncoder,
			// StacktraceKey:  "stacktrace",
		},
	}
	logger, err := cfg.Build(zap.AddCallerSkip(1))
	if err != nil {
		fmt.Println("log build fail:" + err.Error())
	}
	gLog = logger
}

func Debug(msg string, fields ...zap.Field) {
	gLog.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	gLog.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	gLog.Error(msg, fields...)
}

func Destroy() {
	_ = gLog.Sync()
}
