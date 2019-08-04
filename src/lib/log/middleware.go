package log

import (
	"net/http"

	"github.com/abyssparanoia/rapid-go-worker/src/lib/util"
)

// Middleware ... ロガー
type Middleware struct {
	Writer         Writer
	MinOutSeverity Severity
}

// Handle ... ロガーを初期化する
func (m *Middleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// ロガーをContextに設定
		traceID := util.StrUniqueID()
		logger := NewLogger(m.Writer, m.MinOutSeverity, traceID)
		ctx := r.Context()
		ctx = SetLogger(ctx, logger)

		// 実行
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}

// NewMiddleware ... ミドルウェアを作成する
func NewMiddleware(writer Writer, minOutSeverity string) *Middleware {
	mos := NewSeverity(minOutSeverity)
	return &Middleware{
		Writer:         writer,
		MinOutSeverity: mos,
	}
}
