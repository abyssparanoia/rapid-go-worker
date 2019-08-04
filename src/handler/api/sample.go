package api

import (
	"net/http"

	"github.com/abyssparanoia/rapid-go-worker/src/lib/renderer"
	"github.com/abyssparanoia/rapid-go-worker/src/service"
)

// SampleHandler ... サンプルのハンドラ
type SampleHandler struct {
	Svc service.Sample
}

// Sample ... サンプルハンドラ
func (h *SampleHandler) Sample(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	err := h.Svc.Sample(ctx)
	if err != nil {
		renderer.HandleError(ctx, w, "h.Svc.Sample", err)
		return
	}

	renderer.Success(ctx, w)
}

// NewSampleHandler ... ハンドラを作成する
func NewSampleHandler(svc service.Sample) *SampleHandler {
	return &SampleHandler{
		Svc: svc,
	}
}
