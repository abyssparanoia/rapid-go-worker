package log

import (
	"time"
)

// Writer ... ログの出力
type Writer interface {
	Application(
		severity Severity,
		traceID string,
		msg string,
		file string,
		line int64,
		function string,
		at time.Time)
}
