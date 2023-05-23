package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

func main() {
	topic := "vcz"

	w := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	defer func() {
		if err := w.Close(); err != nil {
			log.Fatal("failed to close writer:", err)
		}
	}()

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte("One!"),
		},
		kafka.Message{
			Key:   []byte("Key-B"),
			Value: []byte("Two!"),
		},
		kafka.Message{
			Key:   []byte("Key-C"),
			Value: []byte("Three!"),
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

}
