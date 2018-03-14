// main.go
package main

import (
	"log"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"cloud.google.com/go/pubsub"

	"golang.org/x/net/context"
)

func hello() (payload string, err error) {

	var topic *pubsub.Topic
	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, "terraform-demo-project")
	if err != nil {
		log.Fatal(err)
	}

	topic, _ = client.CreateTopic(ctx, "demo-topic")

	msg := &pubsub.Message{
		Data: []byte(payload),
	}

	if _, err := topic.Publish(ctx, msg).Get(ctx); err != nil {
		return "error", nil
	}

	fmt.Print("Message published.")

	return "Hello!", nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(hello)
}
