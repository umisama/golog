package log

import (
	"fmt"
	"io"
	"runtime"
	"strconv"
	"strings"
	"text/template"
	"time"
)

// LogLevel reprecents log priority.
type LogLevel int

// LogLevel_* reprecents supporting log level.
// priority : Debug < Info < Warn < Critical
// LogLevel_Silent ignores all log levels.
const (
	LogLevel_Debug = LogLevel(iota)	// output all
	LogLevel_Info					// outuut Info or greater.
	LogLevel_Warn					// output Warn and Critical
	LogLevel_Critical				// output Critical only
	LogLevel_Silent					// no output
)

type logger struct {
	log_tmpl *template.Template
	time_fmt string
	level    LogLevel
	prefix   string
	dst      io.Writer
}

type debug_logger struct {
	*logger
}

type info_logger struct {
	*debug_logger
}

type warn_logger struct {
	*info_logger
}

type critical_logger struct {
	*warn_logger
}

type silent_logger struct {
	*critical_logger
}

type Logger interface {
	// Debug output log by Debug level(1)
	Debug(a ...interface{})

	// Debugf output log by Debug level(1) with format string.
	Debugf(msg_fmt string, a ...interface{})

	// Info output log by Info level(2)
	Info(a ...interface{})

	// Infof output log by Info level(2) with format string.
	Infof(msg_fmt string, a ...interface{})

	// Warn output log by Warn level(3)
	Warn(a ...interface{})

	// Warnf output log by Warn level(3) with format string.
	Warnf(msg_fmt string, a ...interface{})

	// Critical output log by Critical level(4)
	Critical(a ...interface{})

	// Criticalf output log by Critical level(4) with format string.
	Criticalf(msg_fmt string, a ...interface{})
}

// LogTemplate is template object on LOG_FORMAT_*
// this object public for documentation.
type LogTemplate struct {
	Time          string
	FuncName      string
	ShortFuncName string
	FileName      string
	ShortFileName string
	LineNumber    string
	Message       string
}

// LOG_FORMAT_* is example format string(with text/template)
const (
	LOG_FORMAT_SIMPLE   = "{{.Time}} : {{.Message}} \n"
	LOG_FORMAT_STANDARD = "{{.Time}} {{.ShortFileName}}:({{.LineNumber}}) : {{.Message}}\n"
	LOG_FORMAT_POWERFUL = "{{.Time}} {{.ShortFileName}}:{{.LineNumber}}({{.ShortFuncName}}) : {{.Message}}\n"
)

// TIME_FORMAT_* is example format string(with time)
const (
	TIME_FORMAT_DATE     = "2006/1/2"
	TIME_FORMAT_SEC      = "2006/1/2 15:04:05"
	TIME_FORMAT_MILLISEC = "2006/1/2 15:04:05.000"
)

// NewLogger returns Logger that outputs to dst with level.  
// log will be formated by time_fmt and log_fmt.
func NewLogger(dst io.Writer, time_fmt string, log_fmt string, level LogLevel) (l Logger, err error) {
	t, err := template.New("log").Parse(log_fmt)
	if err != nil {
		return
	}

	l = &logger{
		level:    level,
		log_tmpl: t,
		time_fmt: time_fmt,
		dst:      dst,
	}

	if level >= LogLevel_Debug {
		l = &debug_logger{l.(*logger)}
	}

	if level >= LogLevel_Info {
		l = &info_logger{l.(*debug_logger)}
	}

	if level >= LogLevel_Warn {
		l = &warn_logger{l.(*info_logger)}
	}

	if level >= LogLevel_Critical {
		l = &critical_logger{l.(*warn_logger)}
	}

	if level >= LogLevel_Silent {
		l = &silent_logger{l.(*critical_logger)}
	}

	return
}

func (l *logger) Debug(a ...interface{}) {
	l.print(a...)
	return
}

func (l *logger) Debugf(msg_fmt string, a ...interface{}) {
	l.printf(msg_fmt, a...)
	return
}

func (l *logger) Info(a ...interface{}) {
	l.print(a...)
	return
}

func (l *logger) Infof(msg_fmt string, a ...interface{}) {
	l.printf(msg_fmt, a...)
	return
}

func (l *logger) Warn(a ...interface{}) {
	l.print(a...)
	return
}

func (l *logger) Warnf(msg_fmt string, a ...interface{}) {
	l.printf(msg_fmt, a...)
	return
}

func (l *logger) Critical(a ...interface{}) {
	l.print(a...)
	return
}

func (l *logger) Criticalf(msg_fmt string, a ...interface{}) {
	l.printf(msg_fmt, a...)
	return
}

func (l *logger) print(a ...interface{}) {
	s := ""
	for _, v := range a {
		s += fmt.Sprintf("%#v ", v)
	}

	l.printer(s)
}

func (l *logger) printf(msg_fmt string, a ...interface{}) {
	s := fmt.Sprintf(msg_fmt, a...)
	l.printer(s)
}

func (l *logger) printer(str string) {
	pc, file_name, line_num, ok := runtime.Caller(3)
	if !ok {
		return
	}

	func_name := runtime.FuncForPC(pc).Name()
	func_name_s := func_name[strings.LastIndex(func_name, ".")+1:]
	file_name_s := file_name[strings.LastIndex(file_name, "/")+1:]

	d := &LogTemplate{
		Time:          time.Now().Format(l.time_fmt),
		FuncName:      func_name,
		ShortFuncName: func_name_s,
		FileName:      file_name,
		ShortFileName: file_name_s,
		LineNumber:    strconv.Itoa(line_num),
		Message:       str,
	}

	l.log_tmpl.Execute(l.dst, d)

	return
}

func (l LogLevel) String() string {
	switch l {
	case LogLevel_Debug:
		return "debug"
	case LogLevel_Info:
		return "info"
	case LogLevel_Warn:
		return "warn"
	case LogLevel_Critical:
		return "critical"
	}

	return "!!panic!!"
}

func (l *info_logger) Debug(a ...interface{}) {
	return
}

func (l *info_logger) Debugf(msg_fmt string, a ...interface{}) {
	return
}

func (l *warn_logger) Info(a ...interface{}) {
	return
}

func (l *warn_logger) Infof(msg_fmt string, a ...interface{}) {
	return
}

func (l *critical_logger) Warn(a ...interface{}) {
	return
}

func (l *critical_logger) Warnf(msg_fmt string, a ...interface{}) {
	return
}

func (l *silent_logger) Critical(a ...interface{}) {
	return
}

func (l *silent_logger) Criticalf(msg_fmt string, a ...interface{}) {
	return
}
