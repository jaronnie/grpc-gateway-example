package logx

import (
	"runtime"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Log output info message
func Log(msg string, fields ...zap.Field) {

	logInstance.Info(msg, fields...)
}

// Debug output debug message
func Debug(msg string, fields ...zap.Field) {
	funcName, caller := stackTrace()
	logInstance.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Debug(msg, fields...)
}

// Debugf output debug message
func Debugf(template string, args ...interface{}) {
	funcName, caller := stackTrace()
	logInstance.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Sugar().Debugf(template, args...)
}

// Debugw output debug message
func Debugw(msg string, keysAndValues ...interface{}) {
	funcName, caller := stackTrace()
	logInstance.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Sugar().Debugw(msg, keysAndValues...)
}

// Info output info message
func Info(msg string, fields ...zap.Field) {
	funcName, caller := stackTrace()
	logInstance.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Info(msg, fields...)
}

// Infof output info message
func Infof(template string, args ...interface{}) {
	funcName, caller := stackTrace()
	logInstance.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Sugar().Infof(template, args...)
}

// Infow output info message
func Infow(msg string, keysAndValues ...interface{}) {
	funcName, caller := stackTrace()
	logInstance.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Sugar().Infow(msg, keysAndValues...)
}

// Warn output warn message
func Warn(msg string, fields ...zap.Field) {
	funcName, caller := stackTrace()
	logInstance.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Warn(msg, fields...)
}

// Warnf output warn message
func Warnf(template string, args ...interface{}) {
	funcName, caller := stackTrace()
	logInstance.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Sugar().Warnf(template, args...)
}

// Warnw output warn message
func Warnw(msg string, keysAndValues ...interface{}) {
	funcName, caller := stackTrace()
	logInstance.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Sugar().Warnw(msg, keysAndValues...)
}

// Error output error message
func Error(msg string, fields ...zap.Field) {
	funcName, caller := stackTrace()
	logInstance.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Error(msg, fields...)
}

// Errorf output error message
func Errorf(template string, args ...interface{}) {
	funcName, caller := stackTrace()
	logInstance.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Sugar().Errorf(template, args...)
}

// Errorw output error message
func Errorw(msg string, keysAndValues ...interface{}) {
	funcName, caller := stackTrace()
	logInstance.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Sugar().Errorw(msg, keysAndValues...)
}

// DPanic output dPanic message
func DPanic(msg string, fields ...zap.Field) {
	funcName, caller := stackTrace()
	logInstance.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).DPanic(msg, fields...)
}

// DPanicf output dPanic message
func DPanicf(template string, args ...interface{}) {
	funcName, caller := stackTrace()
	logInstance.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Sugar().DPanicf(template, args...)
}

// DPanicw output dPanic message
func DPanicw(msg string, keysAndValues ...interface{}) {
	funcName, caller := stackTrace()
	logInstance.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Sugar().DPanicw(msg, keysAndValues...)
}

// Panic output panic message
func Panic(msg string, fields ...zap.Field) {
	funcName, caller := stackTrace()
	logInstance.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Panic(msg, fields...)
}

// Panicf output panic message
func Panicf(template string, args ...interface{}) {
	funcName, caller := stackTrace()
	logInstance.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Sugar().Panicf(template, args...)
}

// Panicw output panic message
func Panicw(msg string, keysAndValues ...interface{}) {
	funcName, caller := stackTrace()
	logInstance.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Sugar().Panicw(msg, keysAndValues...)
}

// Fatal output fatal message
func Fatal(msg string, fields ...zap.Field) {
	funcName, caller := stackTrace()
	logInstance.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Fatal(msg, fields...)
}

// Fatalf output fatal message
func Fatalf(template string, args ...interface{}) {
	funcName, caller := stackTrace()
	logInstance.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Sugar().Fatalf(template, args...)
}

// Fatalw output fatal message
func Fatalw(msg string, keysAndValues ...interface{}) {
	funcName, caller := stackTrace()
	logInstance.With(zap.Any("funcName", funcName)).With(zap.Any("caller", caller)).Sugar().Fatalw(msg, keysAndValues...)
}

// Sugar return zap SugaredLogger instance
func Sugar() *zap.SugaredLogger {
	return logInstance.Sugar()
}

// Named return zap instance
func Named(s string) *zap.Logger {
	return logInstance.Named(s)
}

// WithOptions log with option
func WithOptions(opts ...zap.Option) *zap.Logger {
	return logInstance.WithOptions(opts...)
}

// With log with field
func With(fields ...zap.Field) *zap.Logger {
	return logInstance.With(fields...)
}

// Check level check
func Check(lvl zapcore.Level, msg string) *zapcore.CheckedEntry {
	return logInstance.Check(lvl, msg)
}

// Sync log sync
func Sync() error {
	return logInstance.Sync()
}

// Core return log core
func Core() zapcore.Core {
	return logInstance.Core()
}

func stackTrace() (funcName, caller string) {
	buf := make([]byte, 16384)
	buf = buf[:runtime.Stack(buf, false)]
	stacks := strings.Split(string(buf), "\n")

	funcName = stacks[5]
	funcName = funcName[strings.LastIndex(funcName, "/")+1 : strings.LastIndex(funcName, "(")]
	callers := strings.Split(stacks[6][1:], " ")

	return funcName, callers[0]
}
