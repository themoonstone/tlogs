package tlogs

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"runtime"
	"time"
)

type TLogs struct {
	Entry *logrus.Entry
}

type TLogsHook struct {

}

func NewLogs() *TLogs {
	tlogs := new(TLogs)
	tlogs.Entry = logrus.NewEntry(logrus.New())
	return tlogs
}

func (tlogs_hook *TLogsHook)Fire(entry *logrus.Entry) error {
	file, line, func_name :=  getFileInfo()
	entry.Data["file"] = file
	entry.Data["line"] = line
	entry.Data["func"] = func_name
	return nil
}

func (tlogs_hook *TLogsHook)Levels() []logrus.Level {
	return logrus.AllLevels
}

func (tlogs *TLogs) logStr(kv ...interface{}) string {
	//增加traceId,spanid,pspanid
	if len(kv)%2 != 0 {
		kv = append(kv, "unknown")
	}
	strFmt := ""
	args := []interface{}{}
	for i := 0; i < len(kv); i += 2 {
		strFmt += "{%v:%+v}"
		args = append(args, kv[i], kv[i+1])
	}
	str := fmt.Sprintf(strFmt, args...)
	return str
}

func (tlogs *TLogs) Info(args... interface{})  {
	tlogs.Log(logrus.InfoLevel, tlogs.logStr(args...))
}

func (tlogs *TLogs) Debug(args... interface{})  {
	tlogs.Log(logrus.DebugLevel, tlogs.logStr(args...))
}

func (tlogs *TLogs) Warn(args... interface{})  {
	tlogs.Log(logrus.WarnLevel, tlogs.logStr(args...))
}

func (tlogs *TLogs) Error(args... interface{})  {
	tlogs.Log(logrus.ErrorLevel, tlogs.logStr(args...))
}

func (tlogs *TLogs) Panic(args... interface{})  {
	tlogs.Log(logrus.PanicLevel, tlogs.logStr(args...))
}

func (tlogs *TLogs) Fatal(args... interface{})  {
	tlogs.Log(logrus.FatalLevel, tlogs.logStr(args...))
}

func (tlogs *TLogs) Trace(args... interface{})  {
	tlogs.Log(logrus.TraceLevel, tlogs.logStr(args...))
}
func (tlogs *TLogs) Log(level logrus.Level, args ...interface{})  {
	tlogs.Entry.Log(level, args...)
}

// 设置日志属性
func (tlogs *TLogs) SetAttribute(logs_name string, time_format string, duration time.Duration)  {
	if time_format == "" {
		time_format = default_format
	}
	if duration == 0 {
		duration = default_duration
	}
	if logs_name == "" {
		logs_name = default_dir
	}
	rl, _ := rotatelogs.New(fmt.Sprintf("%s.%s",logs_name, time_format),
		rotatelogs.WithLinkName(logs_name),
		rotatelogs.WithRotationTime(duration))
	tlogs.Entry.Logger.SetOutput(rl)
	tlogs.Entry.Logger.SetFormatter(&logrus.JSONFormatter{})
	tl_hook := &TLogsHook{}
	tlogs.Entry.Logger.AddHook(tl_hook)
}

// 获取文件名、行号、函数调用信息
func getFileInfo() (file string, line int, func_name string) {
	var pc uintptr
	pc, file, line, _ = runtime.Caller(8)
	func_name = runtime.FuncForPC(pc).Name()
	return
}

/*
	1. 设置log全局变量(在main函数中设置)
	2. 每次打印日志时先获取这个全局变量，如果有，直接使用，如果没有，先设置一下
*/

func (tlogs *TLogs) SetTopic(topic string)  {
	// 设置log全局变量(在main函数中设置)
	//tlogs.SetAttribute(default_dir, default_duration)
}