package logging

import (
	"os"
	"fmt"
	"log"
	"time"
)

var (
	// LogSavePath 保存路径
	LogSavePath = "runtime/logs/"
	// LogSaveName 保存文件名
	LogSaveName = "log"
	// LogFileExt 日志文件类型
	LogFileExt = "log"
	// TimeFormat 日志文件名保存格式
	TimeFormat = "20060702"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)

	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func openLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("Permission: %v", err)
	}

	handler, err := os.OpenFile(filePath, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile: %v", err)
	}

	return handler
}

func mkDir() {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir + "/" + getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
