package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/Algatux/queues-experiments/internal"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"time"
)

func main() {

	topicArn := flag.String("topicArn", "", "sns topic arn")
	msgNumber := flag.Int("msgNumber", 1, "number of messages to send (default 1)")
	msgRate := flag.Int("msgRate", 5, "send rate in seconds (default 5)")
	flag.Parse()

	internal.Logger.Debug(fmt.Sprintf("TopicArn: %s", *topicArn))

	cfg := internal.LoadCredentials()

	client := sns.NewFromConfig(cfg)

	for i := *msgNumber; i > 0; i-- {
		publish, err := publishEvent(client, topicArn)
		if err != nil {
			internal.Logger.Error(err)
		}
		internal.Logger.Info(fmt.Sprintf("sent message id `%s`", *publish.MessageId))
		if i > 1 {
			time.Sleep(time.Duration(*msgRate) * time.Second)
		}
	}

}

func publishEvent(client *sns.Client, topicArn *string) (*sns.PublishOutput, error) {
	eventId, msg := internal.GenerateEvent()

	input := &sns.PublishInput{
		Message:                aws.String(msg),
		MessageDeduplicationId: aws.String(eventId),
		MessageGroupId:         aws.String("events"),
		TopicArn:               topicArn,
	}

	return client.Publish(context.Background(), input)
}
