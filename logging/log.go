package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// Level 等级
type Level int

var (
	// F 文件句柄
	F *os.File

	// DefaultPrefix 前缀
	DefaultPrefix = ""
	// DefaultCallerDepth 层级
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	// DEBUG 调试
	DEBUG Level = iota
	// INFO 提醒
	INFO
	// WARNING 警告
	WARNING
	// ERROR 错误
	ERROR
	// FATAL 致命
	FATAL
)

func init() {
	filePath := getLogFileFullPath()
	F = openLogFile(filePath)

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

// Debug 开发
func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}

// Info 提醒
func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v)
}

// Warn 警告
func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v)
}

// Error 错误
func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v)
}

// Fatal 致命
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v)
}

func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}
