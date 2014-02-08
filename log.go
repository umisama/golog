package log

import (
	"fmt"
	"io"
	"runtime"
	"strconv"
	"text/template"
	"time"
)

type LogLevel int

const (
	LogLevel_Debug = LogLevel(iota)
	LogLevel_Info
	LogLevel_Warn
	LogLevel_Critical
)

type Logger struct {
	log_tmpl *template.Template
	time_fmt string
	level    int64
	prefix   string
	dst      io.Writer
}

type LogTemplate struct {
	Time          string
	FuncName      string
	FileName      string
	LineNumber    string
	Message       string
}

const (
	LOG_FORMAT_SIMPLE   = "{{.Time}} : {{.Message}} \n"
	LOG_FORMAT_STANDARD = "{{.Time}} {{.FileName}}:({{.LineNumber}}) : {{.Message}}\n"
	LOG_FORMAT_POWERFUL = "{{.Time}} {{.FileName}}:{{.LineNumber}}({{.FuncName}}) : {{.Message}}\n"
)

const (
	TIME_FORMAT_DATE = "2006/1/2"
	TIME_FORMAT_SEC = "2006/1/2 15:04:05"
	TIME_FORMAT_MILLISEC = "2006/1/2 15:04:05.000"
)

func NewLogger(dst io.Writer, time_fmt string, log_fmt string) (logger *Logger, err error) {
	t, err := template.New("log").Parse(log_fmt)
	if err != nil {
		return
	}

	logger = &Logger{
		log_tmpl: t,
		time_fmt: time_fmt,
		level:    10,
		dst:      dst,
	}

	return
}

func NewSimpleLogger(dst io.Writer) *Logger {
	logger := new(Logger)

	return logger
}

func (l *Logger) SetEnableLevel(level int64) {
	l.level = level
	return
}

func (l *Logger) Print(level int64, a ...interface{}) {
	if l.level < level {
		return
	}

	s := ""
	for _, v := range a {
		s += fmt.Sprintf("%#v ", v)
	}

	l.printer(s)
}

func (l *Logger) Printf(level int64, msg_fmt string, a ...interface{}) {
	s := fmt.Sprintf(msg_fmt, a...)
	l.printer(s)
}

func (l *Logger) printer(str string) {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return
	}

	go func() {
		d := &LogTemplate{
			Time:       time.Now().Format(l.time_fmt),
			FuncName:   runtime.FuncForPC(pc).Name(),
			FileName:   file,
			LineNumber: strconv.Itoa(line),
			Message:    str,
		}

		l.log_tmpl.Execute(l.dst, d)
	}()

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
