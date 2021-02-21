package logger

import (
	"io"
	"time"

	"github.com/go-kit/kit/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger
var logLevel zapcore.Level
var atom = zap.NewAtomicLevel()

// InitLog init
func InitLog(logFile, level string) {
	logLevel = getLogLevel(level)
	logger = NewLogger(logFile, level, true).Logger
}

// Logger struct
type Logger struct {
	*zap.Logger
	io.Closer
}

// NewLogger @param defaultAtom :if use the default atom
// create by tb
// use this to create logger for different log file
// compatible with old one
// use the global  variable atom,be care
func NewLogger(logFile string, level string, defaultAtom bool) Logger {
	f := &lumberjack.Logger{
		Filename:   logFile, // 日志文件路径
		MaxSize:    1024,    // megabytes
		MaxBackups: 3,       // 最多保留3个备份
		MaxAge:     7,       // 最多保存days
		Compress:   true,    // 是否压缩 disabled by default
	}
	writer := zapcore.AddSync(f)

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "T",
		LevelKey:      "L",
		NameKey:       "N",
		CallerKey:     "C",
		MessageKey:    "M",
		StacktraceKey: "S",

		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     timeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	at := atom
	if defaultAtom {
		// use the global log level
		atom.SetLevel(getLogLevel(level))
	} else {
		at := zap.NewAtomicLevel()
		at.SetLevel(getLogLevel(level))
	}
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		writer,
		at, //level,
	)

	caller := zap.AddCaller()
	logger := Logger{
		Logger: zap.New(core, caller, zap.AddCallerSkip(1)),
		Closer: f,
	}
	return logger
}

// func funcCaller(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
// 	enc.AppendString(filepath.Base(caller.FullPath()))
// }
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000 -0700"))
}

func getLogLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}

// Fatal 根据模式
func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}

// Error print
func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

// Warn print
func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

// Debug print
func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

// Info print
func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

// SetLogLevel set log level
func SetLogLevel(level string) {
	logLevel = getLogLevel(level)
	atom.SetLevel(logLevel)
}

// GetLogger get looger
func GetLogger() *zap.Logger {
	return logger
}

type kitLogger func(msg string, keysAndValues ...interface{})

func (l kitLogger) Log(kv ...interface{}) error {
	l("kit_msg", kv...)
	return nil
}

// GetKitLogger returns a Go kit log.Logger that sends
// log events to a zap.Logger.
func GetKitLogger() log.Logger {
	sugarLogger := logger.WithOptions(zap.AddCallerSkip(1)).Sugar()
	var sugar kitLogger
	switch logLevel {
	case zapcore.DebugLevel:
		sugar = sugarLogger.Debugw
	case zapcore.InfoLevel:
		sugar = sugarLogger.Infow
	case zapcore.WarnLevel:
		sugar = sugarLogger.Warnw
	case zapcore.ErrorLevel:
		sugar = sugarLogger.Errorw
	case zapcore.DPanicLevel:
		sugar = sugarLogger.DPanicw
	case zapcore.PanicLevel:
		sugar = sugarLogger.Panicw
	case zapcore.FatalLevel:
		sugar = sugarLogger.Fatalw
	default:
		sugar = sugarLogger.Infow
	}
	return sugar
}
