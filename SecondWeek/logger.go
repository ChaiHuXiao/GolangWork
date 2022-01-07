package main

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"time"
)

const (
	INFO    = 0
	DEBUG   = 1
	TRACE   = 2
	WARNING = 3
	ERROR   = 4
	FATAL   = 5
)

type Conf struct {
	Logconf Logconf
}

type Logconf struct {
	Path string
	Leve int8
}

// 按照日期存储日志信息
type Logger2 struct {
	level    int8       // 打印日志级别
	path     string     // 文件存储路径
	filename string     // 文件名称
	file     *os.File   //文件指针
	time     *time.Time // 文件日期
}

// 构造方法
func NewLogger2(path string) Logger2 {
	l := Logger2{path: path}
	l.newLogFile()
	return l
}

// 设置日志级别
func (this *Logger2) SetLevel(level int8) {
	this.level = level
}

// 设置日志文件路径
func (this *Logger2) SetPath(path string) {
	this.path = path
}

// 向日志中追加内容
func (this *Logger2) writeToLog(level int8, msg string) {
	if this == nil {
		println("对象创建失败")
	}

	// 当前日志文件已过期，需要用新的日志文件记录，更新Logger信息
	if !this.isOneDay() {
		this = this.newLogFile()
	}
	var l string
	switch level {
	case DEBUG:
		l = "DEBUG"
	case TRACE:
		l = "TRACE"
	case INFO:
		l = "INFO"
	case WARNING:
		l = "WARNING"
	case ERROR:
		l = "ERROR"
	case FATAL:
		l = "FATAL"
	}
	// 当前级别大于日志级别才输出 否则不输出
	if level >= this.level {
		_, file, line, ok := runtime.Caller(2)
		if !ok {
			fmt.Fprintln(this.file, "[", time.Now().Format("2006-01-02 15:04:05"), "]", "[ERROR]", "[", file, ":", line, "]", "runtime.Caller() fail")
			return
		}
		// 日志信息写入文件中
		fmt.Fprintln(this.file, "[", time.Now().Format("2006-01-02 15:04:05"), "]", "[", l, "]", "[", file, ":", line, "]", msg)
	}
}

// 实现Log接口中的方法
func (this Logger2) Debug(msg string) {
	this.writeToLog(DEBUG, msg)
}
func (this Logger2) TRACE(msg string) {
	this.writeToLog(TRACE, msg)
}
func (this Logger2) INFO(msg string) {
	this.writeToLog(INFO, msg)
}
func (this Logger2) WARNING(msg string) {
	this.writeToLog(WARNING, msg)
}
func (this Logger2) ERROR(msg string) {
	this.writeToLog(ERROR, msg)
}
func (this Logger2) FATAL(msg string) {
	this.writeToLog(FATAL, msg)
}

// 判断当前time是否是同一天
func (this *Logger2) isOneDay() bool {
	println("当前级别为：")
	println(this.level)
	now := time.Now()
	day := this.time
	println(this.time.Day())
	println(time.Now().Month())
	println(now.Day() == day.Day())
	println(now.Month() == day.Month())
	println(now.Year() == day.Year())
	// 比较年月日是否相同
	return now.Day() == day.Day() && now.Month() == day.Month() && now.Year() == day.Year()
}

// 若当前是新的一天，则需要创建新文件，同时更新文件信息
func (this *Logger2) newLogFile() *Logger2 {
	// 获取当前日期
	now := time.Now()
	filename := now.Format("20060102") + ".log"
	// 获取日志路径
	road := this.path
	newPath := path.Join(road, filename)
	print(newPath)
	// 创建新的文件 以当前年月日命名
	file, err := os.OpenFile(newPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		// 将err作为日志输出在文件中
		_, f, line, _ := runtime.Caller(0)
		fmt.Fprintln(this.file, "[", time.Now().Format("2006-01-02 15:04:05"), "]", "[ERROR]", "[", f, ":", line, "]", "create log file failed!")
		return this
	}
	// 更新结构体数据
	this.filename = filename
	this.file = file
	this.time = &now
	return this
}
