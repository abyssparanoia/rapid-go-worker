package cloudpubsub

import (
	"context"

	pubsub "cloud.google.com/go/pubsub"
	"github.com/abyssparanoia/rapid-go-worker/src/lib/log"
	"github.com/abyssparanoia/rapid-go-worker/src/lib/util"
	"google.golang.org/api/option"
)

type ps struct {
	subscription   *pubsub.Subscription
	writer         log.Writer
	minOutSeverity log.Severity
}

func (p *ps) Listen(handle MessageHandler) error {

	ctx := context.Background()

	// ロガーをContextに設定
	traceID := util.StrUniqueID()
	logger := log.NewLogger(p.writer, p.minOutSeverity, traceID)
	ctx = log.SetLogger(ctx, logger)

	log.Debugf(ctx, "start subscription")

	err := p.subscription.Receive(ctx, handle)

	if err != nil {
		return err
	}

	return nil
}

// NewSubscription ... get subscriprio
func NewSubscription(projectID string, credentialsPath string, subscriptionID string, writer log.Writer, minOutSeverity log.Severity) Subscription {
	ctx := context.Background()
	opt := option.WithCredentialsFile(credentialsPath)
	psClient, err := pubsub.NewClient(ctx, projectID, opt)
	if err != nil {
		panic(err)
	}

	subscription := psClient.Subscription(subscriptionID)

	return &ps{
		subscription: subscription,
	}
}
