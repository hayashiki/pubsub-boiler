package main

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"pubsub-boiler"
)

func main()  {
	var message string
	topic := "event-fired"
	ctx := context.Background()
	proj := os.Getenv("GCP_PROJECT")

	rootCmd := &cobra.Command{
		Use:   "sub",
		Short: "subscriber",
		Long:  "subscriber",
		RunE: func(cmd *cobra.Command, args []string) error {
			sub, err := pubsubboiler.NewSubscription(ctx, proj)
			if err = sub.Subscribe(ctx, topic); err != nil {
				fmt.Printf("err %v", err)
				return err
			}
			return nil
		},
	}
	publishCmd := &cobra.Command{
		Use:   "publish",
		Short: "Publish message",
		Long:  `Publish message`,
		RunE: func(cmd *cobra.Command, args []string) error {
			pub, err := pubsubboiler.NewPublisher(ctx, proj)
			if err != nil {
				return err
			}
			return pub.Publish(ctx, topic, []byte(message))
		},
	}
	publishCmd.Flags().StringVarP(&message, "message", "m", "hello world", "publish message (default is hello world)")

	rootCmd.AddCommand(publishCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
