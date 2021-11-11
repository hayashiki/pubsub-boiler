package pubsubboiler

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/pubsub"
)

type Subscription struct {
	client *pubsub.Client
}

func NewSubscription(ctx context.Context, projectID string) (*Subscription, error) {
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}

	subscription := Subscription{
		client: client,
	}

	return &subscription, nil
}

func (p *Subscription) FindSubscription(ctx context.Context, topicName string) (*pubsub.Subscription, error) {
	it := p.client.Subscriptions(ctx)

	for {
		subscription, err := it.Next()
		if err != nil {
			return nil, nil
		}
		if subscription.ID() == topicName {
			return subscription, nil
		}
	}

	return nil, nil
}

func (p *Subscription) CreateSubscription(ctx context.Context, topicName string) (*pubsub.Subscription, error) {
	topic, err := initializePubsubTopic(ctx, p.client, topicName)

	if err != nil {
		return nil, fmt.Errorf("could not open topic subscription: %v", err)
	}

	sub, err := p.client.CreateSubscription(ctx, topicName, pubsub.SubscriptionConfig{
		Topic:       topic,
		AckDeadline: 10 * time.Second,
	})

	return sub, err
}

func (p *Subscription) Subscribe(ctx context.Context, topicName string) error {
	var sub *pubsub.Subscription
	sub, err := p.FindSubscription(ctx, topicName)
	if err != nil {
		return fmt.Errorf("could not find subscription: %v", err)
	}
	if sub == nil {
		sub, err = p.CreateSubscription(ctx, topicName)
		if err != nil {
			return fmt.Errorf("could not create subscription: %v", err)
		}
	}
	err = sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		log.Printf("Got message: %s", m.Data)
		m.Ack()
	})
	if err != nil {
		log.Printf("Got message: %s", err.Error())
	}
	return err
}
