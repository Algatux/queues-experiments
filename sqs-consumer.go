package main

import (
	"context"
	"flag"
	"github.com/Algatux/queues-experiments/internal"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func main() {

	endpoint := flag.String("endpoint", "", "sqs endpoint")
	flag.Parse()
	internal.Logger.Debug(*endpoint)

	cfg := internal.LoadCredentials()
	client := sqs.NewFromConfig(cfg)

	for {
		receiveParams := &sqs.ReceiveMessageInput{
			QueueUrl:            endpoint,
			MaxNumberOfMessages: 10,
			WaitTimeSeconds:     10,
			VisibilityTimeout:   30,
		}

		ctx := context.Background()
		resp, err := client.ReceiveMessage(ctx, receiveParams)
		if err != nil {
			internal.Logger.FatalError("Errore durante la ricezione dei messaggi:", err)
		}

		if len(resp.Messages) == 0 {
			internal.Logger.Info("Nessun messaggio ricevuto.")
			continue
		}

		for _, message := range resp.Messages {
			internal.Logger.Info("Messaggio ricevuto:", aws.ToString(message.Body))
			go func() {
				deleteParams := &sqs.DeleteMessageInput{
					QueueUrl:      endpoint,
					ReceiptHandle: message.ReceiptHandle,
				}
				_, err := client.DeleteMessage(ctx, deleteParams)
				if err != nil {
					internal.Logger.Error("Errore durante la cancellazione del messaggio:", err)
				} else {
					internal.Logger.Info("Messaggio cancellato con successo.")
				}
			}()
		}
	}

}
