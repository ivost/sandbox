package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"cloud.google.com/go/pubsub"
)

// topic of type https://godoc.org/cloud.google.com/go/pubsub#Topic

const (
	projectID = "robotic-energy-245402"
	subName   = projectID + "/my-subscription"
	topic     = "topic1"
)

//? const sub = "projects/robotic-energy-245402/subscriptions/my-subscription"

func main() {
	pullMsgs(projectID, subName)
}

func pullMsgs(projectID, subName string /*, topic *pubsub.Topic */) error {
	log.Printf("pubsub subscription: %v", subName)
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}

	//// Publish 10 messages on the topic.
	//var results []*pubsub.PublishResult
	//for i := 0; i < 10; i++ {
	//	res := topic.Publish(ctx, &pubsub.Message{
	//		Data: []byte(fmt.Sprintf("hello world #%d", i)),
	//	})
	//	results = append(results, res)
	//}
	//
	//// Check that all messages were published.
	//for _, r := range results {
	//	_, err := r.Get(ctx)
	//	if err != nil {
	//		return fmt.Errorf("Get: %v", err)
	//	}
	//}

	// Consume 10 messages.
	var mu sync.Mutex
	received := 0
	sub := client.Subscription(subName)
	cctx, cancel := context.WithCancel(ctx)
	err = sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		msg.Ack()
		log.Printf("Got message: %q\n", string(msg.Data))
		mu.Lock()
		defer mu.Unlock()
		received++
		if received == 10 {
			cancel()
		}
	})
	if err != nil {
		return fmt.Errorf("receive error: %v", err)
	}
	return nil
}
