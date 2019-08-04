package log

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type writerStackdriver struct {
	ProjectID string
}

func (w *writerStackdriver) Application(
	severity Severity,
	traceID string,
	msg string,
	file string,
	line int64,
	function string,
	at time.Time) {
	e := &Entry{
		Severity: severity.String(),
		Time:     Time(at),
		Trace:    fmt.Sprintf("projects/%s/traces/%s", w.ProjectID, traceID),
		Message:  fmt.Sprintf("%s:%d [%s] %s", file, line, function, msg),
	}
	b, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, string(b)+"\n")
}

// NewWriterStackdriver ... ログ出力を作成する
func NewWriterStackdriver(projectID string) Writer {
	return &writerStackdriver{
		ProjectID: projectID,
	}
}
