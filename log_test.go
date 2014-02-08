package log

import (
	"os"
	"testing"
)

func TestLogPrintf ( t *testing.T ) {
	l, err := NewLogger( os.Stdout, TIME_FORMAT_SEC, LOG_FORMAT_POWERFUL )
	if err != nil {
		t.Errorf( err.Error() )
	}

	l.Printf( 10, "%d", 1024 )
	return
}

func TestLogPrint( t *testing.T ) {
	l, err := NewLogger( os.Stdout, TIME_FORMAT_SEC, LOG_FORMAT_POWERFUL )
	if err != nil {
		t.Errorf( err.Error() )
	}

	l.Print( 10, 1024, "abc", l )
	return
}
