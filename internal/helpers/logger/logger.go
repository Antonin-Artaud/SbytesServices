package logger

import (
	"log"
	"os"
	"runtime"
	"time"
)

type logLevel int

const (
	ERROR logLevel = iota
	WARNING
	INFO
	DEBUG
)

type Logger struct {
	*log.Logger
	level logLevel
}

func NewLogger(level logLevel) *Logger {
	return &Logger{log.New(os.Stdout, "", log.LstdFlags), level}
}

func (l *Logger) LogError(message string) {
	if l.level >= ERROR {
		timeFormat := time.Now().Format("2006-01-02 15:04:05")
		pc, _, _, _ := runtime.Caller(1)
		funcName := runtime.FuncForPC(pc).Name()
		l.Printf("[ERROR] - %s : %s -> %s\n", timeFormat, funcName, message)
	}
}

func (l *Logger) LogWarning(message string) {
	if l.level >= WARNING {
		timeFormat := time.Now().Format("2006-01-02 15:04:05")
		pc, _, _, _ := runtime.Caller(1)
		funcName := runtime.FuncForPC(pc).Name()
		l.Printf("[WARNING] - %s : %s -> %s\n", timeFormat, funcName, message)
	}
}

func (l *Logger) LogInfo(message string) {
	if l.level >= INFO {
		timeFormat := time.Now().Format("2006-01-02 15:04:05")
		pc, _, _, _ := runtime.Caller(1)
		funcName := runtime.FuncForPC(pc).Name()
		l.Printf("[INFO] - %s : %s -> %s\n", timeFormat, funcName, message)
	}
}

func (l *Logger) LogDebug(message string) {
	if l.level >= DEBUG {
		timeFormat := time.Now().Format("2006-01-02 15:04:05")
		pc, _, _, _ := runtime.Caller(1)
		funcName := runtime.FuncForPC(pc).Name()
		l.Printf("[DEBUG] - %s : %s -> %s\n", timeFormat, funcName, message)
	}
}
