package cloudpubsub

import (
	"context"

	pubsub "cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

type ps struct {
	subscription *pubsub.Subscription
}

func (p *ps) Listen(*MessageHandler) error {
	return nil
}

// NewSubscription ... get subscriprio
func NewSubscription(projectID string, credentialsPath string, subscriptionID string) Subscription {
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