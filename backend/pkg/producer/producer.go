package producer

import (
	"context"

	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(cfg ProducerConfig) *Producer {
	return &Producer{writer: &kafka.Writer{
		Addr: kafka.TCP(cfg.Addr),
	}}
}

func (p *Producer) Publish(ctx context.Context, event proto.Message, topic string, key string) error {
	eventBytes, err := protojson.Marshal(event)
	if err != nil {
		return err
	}

	if err = p.writer.WriteMessages(ctx, kafka.Message{
		Topic: topic,
		Key:   []byte(key),
		Value: eventBytes,
	}); err != nil {
		return err
	}

	return nil
}
