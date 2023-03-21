package consumer

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
}

func NewConsumer(cfg ConsumerConfig, topic string) *Consumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{cfg.Addr},
		Topic:    topic,
		MinBytes: 1,
		MaxBytes: 10e6,
	})
	err := reader.SetOffset(kafka.LastOffset)
	if err != nil {
		log.Fatal(err)
	}

	return &Consumer{
		reader: reader,
	}
}

func (c *Consumer) Consume(ctx context.Context) ([]byte, error) {
	message, err := c.reader.ReadMessage(ctx)
	if err != nil {
		return nil, err
	}

	return message.Value, nil
}
