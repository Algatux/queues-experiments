# queues-experiments

### Usage
    go run sqs-consumer.go --endpoint "https://sqs.eu-west-1.amazonaws.com/***/queue-1"

    go run sns-notifier.go --topicArn "arn:aws:sns:eu-west-1:***:event-collector.fifo" --msgNumber 20 --msgRate 1