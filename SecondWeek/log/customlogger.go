package main

import (
	"bytes"
	"io"
	"os"
	"time"
)

const (
	INFO           = 0
	DEBUG          = 1
	WARNING        = 2
	ERROR          = 3
	FILEWRITEERROR = "File write failure "
)

type MyLogger struct {
	level    int8
	filePath string
	fileName string
	file     *os.File
	time     *time.Time
}

func InitializeLogger(path string) MyLogger {
	logger := MyLogger{filePath: path}
	logger.CreateLogFile()
	return logger
}

func (logger *MyLogger) CreateLogFile() *MyLogger {
	now := time.Now()
	today := now.Format("20060102")
	println(today)
	fileName := StringSplicing(12, today, ".log")
	println(fileName)
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend)

	if err != nil {
		panic(time.Now().Format("2006-01-02 15:04:05") + "  ERROR  Failed to create a log file: " + fileName)
	}
	logger.file = file
	logger.fileName = fileName
	logger.time = &now
	return logger
}

func StringSplicing(cap int, joins ...string) string {
	buffer := bytes.Buffer{}
	buffer.Grow(cap)
	for _, v := range joins {
		buffer.WriteString(v)
	}
	return buffer.String()
}

func (logger *MyLogger) Istoday() bool {
	return time.Now().Year() == logger.time.Year() && time.Now().Month() == logger.time.Month() && time.Now().Day() == logger.time.Day()
}

func (logger *MyLogger) writeToLog(msg string) {

	defer logger.file.Close()

	if logger == nil {
		panic("Description Failed to create a log object！")
	}

	if !logger.Istoday() {
		logger = logger.CreateLogFile()
	}

	var curLevel string
	switch logger.level {
	case INFO:
		curLevel = "INFO"
	case DEBUG:
		curLevel = "DEBUG"
	case WARNING:
		curLevel = "WARNING"
	case ERROR:
		curLevel = "ERROR"
	}
	_, err := io.WriteString(logger.file, time.Now().Format("2006-01-02 15:04:05")+curLevel+"  "+msg)
	if err != nil {
		panic(err.Error())
	}

}
