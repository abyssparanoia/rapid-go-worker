package log

import (
	"fmt"
	"time"
)

type writerStdout struct {
	TimeFormat string
}

func (w *writerStdout) Application(
	severity Severity,
	traceID string,
	msg string,
	file string,
	line int64,
	function string,
	at time.Time) {
	date := at.Format(w.TimeFormat)
	fmt.Printf("%s [%s] %s:%d [%s] %s\n", date, severity.String(), file, line, function, msg)
}

// NewWriterStdout ... ログ出力を作成する
func NewWriterStdout() Writer {
	return &writerStdout{
		TimeFormat: "2006-01-02 15:04:05.000",
	}
}
