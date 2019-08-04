package handler

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
	"github.com/abyssparanoia/rapid-go-worker/src/lib/log"
	"github.com/abyssparanoia/rapid-go-worker/src/service"
)

// MessageHandler ... handler
type MessageHandler struct {
	svc service.Sample
}

func (h *MessageHandler) Handle(ctx context.Context, msg *pubsub.Message) {
	msg.Ack()
	fmt.Printf("Got message: %q\n", string(msg.Data))
	log.Debugf(ctx, "Got message: %q\n", string(msg.Data))

}

// NewHandler ... get new handler
func NewHandler(svc service.Sample) *MessageHandler {
	return &MessageHandler{
		svc: svc,
	}
}
