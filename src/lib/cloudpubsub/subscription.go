package cloudpubsub

import (
	"context"

	pubsub "cloud.google.com/go/pubsub"
)

// MessageHandler ... message handler
type MessageHandler func(context.Context, *pubsub.Message)

// Subscription ... subscription interface
type Subscription interface {
	Listen(handle MessageHandler) error
}
