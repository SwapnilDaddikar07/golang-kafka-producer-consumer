package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func main() {
	topic := "my-topic"
	conn, err := getConnection(topic)
	if err != nil {
		log.Fatalf("Could not establish connection to kafka broker %v", err)
		return
	}
	defer func(){
		if err := conn.Close(); err != nil {
			log.Fatal("failed to close writer:", err)
		}
	}()

	log.Printf("Starting a consumer to publish kafka messages every 3 seconds to kafka topic %s \n", topic)

	for {
		bytesWritten, writeErr := conn.WriteMessages(
			kafka.Message{Value: []byte("xyz-try")},
		)
		if writeErr != nil {
			log.Fatalf("could not write to kafka broker %v", writeErr)
			break
		}
		log.Printf("number of bytes written - %d \n", bytesWritten)
		time.Sleep(time.Second * 5)
	}

}

func getConnection(topic string) (*kafka.Conn, error) {
	return kafka.DialLeader(context.Background(), "tcp", "localhost:9093", topic, 2)
}
