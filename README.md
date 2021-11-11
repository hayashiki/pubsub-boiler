# pubsub-boiler

Using Cloud PubSub with the emulator

# start subscribe message

`$ make sub`

# publish message

`$docker-compose exec app go run cmd/sub/main.go publish -m $message`
