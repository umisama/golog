package log

import (
	"bytes"
	"testing"
)

type TestCase struct {
	level   LogLevel
	content []interface{}
	output  string
}

var testformat = "{{.Message}}"

var testcases = []TestCase{
	TestCase{
		level:   LogLevel_Debug,
		content: []interface{}{"hello_debug", "hoyo", "fuga"},
		output:  `"hello_debug" "hoyo" "fuga" `,
	},
	TestCase{
		level: LogLevel_Info,
		content: []interface{}{map[string]string{
			"hello_info":  "hogehoge",
			"hello_info2": "piyopiyo",
		}},
		output: `map[string]string{"hello_info":"hogehoge", "hello_info2":"piyopiyo"} `,
	},
	TestCase{
		level: LogLevel_Warn,
		content: []interface{}{
			struct{ hoge string }{hoge: "hello_warn"},
		},
		output: `struct { hoge string }{hoge:"hello_warn"} `,
	},
	TestCase{
		level:   LogLevel_Critical,
		content: []interface{}{1, 2, 3, 4, 5},
		output:  "1 2 3 4 5 ",
	},
	TestCase{
		level:   LogLevel_Silent,
		content: []interface{}{"silent"},
		output:  "",
	},
}

func TestDebug(t *testing.T) {
	for _, v := range testcases {
		buf := new(bytes.Buffer)
		l, err := NewLogger(buf, TIME_FORMAT_SEC, testformat, v.level)
		if err != nil {
			t.Errorf(err.Error())
		}

		l.Debug( v.content...)
		if v.level <= LogLevel_Debug {
			if buf.String() != v.output {
				t.Errorf("do not match\n(%s)\n(%s)\n", v.output, buf.String())
			}
		} else {
			if buf.String() != "" {
				t.Errorf("unexpected output\n(%s)", buf.String())
			}
		}
	}

	return
}

func TestInfo(t *testing.T) {
	for _, v := range testcases {
		buf := new(bytes.Buffer)
		l, err := NewLogger(buf, TIME_FORMAT_SEC, testformat, v.level)
		if err != nil {
			t.Errorf(err.Error())
		}

		l.Info( v.content...)
		if v.level <= LogLevel_Info{
			if buf.String() != v.output {
				t.Errorf("do not match\n(%s)\n(%s)\n", v.output, buf.String())
			}
		} else {
			if buf.String() != "" {
				t.Errorf("unexpected output\n(%s)", buf.String())
			}
		}
	}

	return
}

func TestWarn(t *testing.T) {
	for _, v := range testcases {
		buf := new(bytes.Buffer)
		l, err := NewLogger(buf, TIME_FORMAT_SEC, testformat, v.level)
		if err != nil {
			t.Errorf(err.Error())
		}

		l.Warn( v.content...)
		if v.level <= LogLevel_Warn{
			if buf.String() != v.output {
				t.Errorf("do not match\n(%s)\n(%s)\n", v.output, buf.String())
			}
		} else {
			if buf.String() != "" {
				t.Errorf("unexpected output\n(%s)", buf.String())
			}
		}
	}

	return
}

func TestCritical(t *testing.T) {
	for _, v := range testcases {
		buf := new(bytes.Buffer)
		l, err := NewLogger(buf, TIME_FORMAT_SEC, testformat, v.level)
		if err != nil {
			t.Errorf(err.Error())
		}

		l.Critical( v.content...)
		if v.level <= LogLevel_Critical{
			if buf.String() != v.output {
				t.Errorf("do not match\n(%s)\n(%s)\n", v.output, buf.String())
			}
		} else {
			if buf.String() != "" {
				t.Errorf("unexpected output\n(%s)", buf.String())
			}
		}
	}

	return
}
