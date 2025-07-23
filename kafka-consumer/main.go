package main

import (
	"context"
	"fmt"
	"log"
	_ "time"

	"github.com/twmb/franz-go/pkg/kgo"
)

func main() {
	brokers := []string{"localhost:9093"}
	topics := []string{"todo-user", "todo-task"}

	client, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
		kgo.ConsumerGroup("todo-consumer-group"),
		kgo.ConsumeTopics(topics...),
	)
	if err != nil {
		log.Fatalf("cannot create Kafka client: %v", err)
	}
	defer client.Close()

	fmt.Println("Consumer is running and listening to topics...")

	for {
		fetches := client.PollFetches(context.Background())

		if errs := fetches.Errors(); len(errs) > 0 {
			for _, err := range errs {
				log.Printf("Kafka fetch error: %v", err)
			}
			continue
		}

		fetches.EachRecord(func(record *kgo.Record) {
			fmt.Printf(
				"Topic: %s | Partition: %d | Offset: %d\n%s\n\n",
				record.Topic, record.Partition, record.Offset,
				string(record.Value),
			)
		})
	}
}

//
