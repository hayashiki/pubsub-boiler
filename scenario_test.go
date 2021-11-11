package pubsubboiler

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestPubSubPublisher_Publish(t *testing.T) {
	topic := "sample.event-fired"
	ctx := context.Background()
	proj := "local"
	os.Setenv("PUBSUB_EMULATOR_HOST", "localhost:8085")
	os.Setenv("PUBSUB_PROJECT_ID", "local")

	pub, err := NewPublisher(ctx, proj)
	if err != nil {
		t.Error(err)
	}
	pub.Publish(ctx, topic, []byte("Hello World"))
	if err != nil {
		t.Error(err)
	}
	sub, err := NewSubscription(ctx, proj)

	go func() {
		if err = sub.Subscribe(ctx, topic); err != nil {
			fmt.Printf("err %v", err)
			t.Error(err)
		}
	}()

	time.Sleep(1 * time.Second)
}
