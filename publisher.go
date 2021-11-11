package pubsubboiler

import (
	"context"
	"log"

	"cloud.google.com/go/pubsub"
)

type Publisher struct {
	client *pubsub.Client
}

func NewPublisher(ctx context.Context, projectID string) (*Publisher, error) {
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}

	return &Publisher{client: client}, nil
}

func (p *Publisher) Publish(ctx context.Context, topicName string, data []byte) error {
	topic, err := initializePubsubTopic(ctx, p.client, topicName)
	if err != nil {
		return err
	}
	msg := pubsub.Message{Data: data}

	res, err := topic.Publish(ctx, &msg).Get(ctx)
	log.Printf("published id:%v", res)
	return err
}
