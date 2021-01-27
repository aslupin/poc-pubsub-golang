package main

import (
	"fmt"
	"log"
	"time"
	"context"
	"cloud.google.com/go/pubsub"
)

func main() {

	pubsubClient := "poc-pubsub-golang"
	topicName := "poc-topic1"
	subName := "poc-sub1"
	msg := "hello world"

	// Create Context and new Pubsub client
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, pubsubClient)

	if err != nil {
		log.Fatal(err)
	}

	// Create topic	
	topic, _  := client.CreateTopic(ctx, topicName)
	// Create a new subscription to the previously
	// created topic and ensure it never expires.
	_, err = client.CreateSubscription(ctx, subName, pubsub.SubscriptionConfig{
		Topic:            topic,
		AckDeadline:      10 * time.Second,
		ExpirationPolicy: time.Duration(0),
	})

	if err != nil {
		log.Fatal(err)
	}

	// If topic is exist use the line
	// topic := client.Topic(topicName)

	// Publish "hello world" on topic1.
	res := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})
	
	// The publish happens asynchronously.
	// Later, you can get the result from res:
	msgID, err := res.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("msgID: ", msgID)

	var data string
	// Use a callback to receive messages via subscription1.
	sub := client.Subscription(subName)
	err = sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		fmt.Println(string(m.Data), data)	
		m.Ack() // Acknowledge that we've consumed the message.
	})
	if err != nil {
		log.Println(err)
	}	
}
