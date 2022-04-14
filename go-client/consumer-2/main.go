package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func main() {

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9093"},
		GroupID: "consumer-group-2",
		Topic:   "my-topic",
	})

	defer func() {
		reader.Close()
	}()

	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("error when consuming message %v \n", err)
			break
		}
		fmt.Printf("consumer-group-2 message value is %s and partition read from is %d \n ", string(message.Value), message.Partition)
		time.Sleep(time.Second * 2)
	}
}
