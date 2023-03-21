package consumer

import (
	"context"
	"log"
	"reflect"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type ConsumerRegistry struct {
	eventProcessors map[string]reflect.Value
	eventConsumers  map[string]*Consumer
}

func NewConsumerRegistry() *ConsumerRegistry {
	return &ConsumerRegistry{
		eventProcessors: make(map[string]reflect.Value),
		eventConsumers:  make(map[string]*Consumer),
	}
}

func (r *ConsumerRegistry) Register(consumer interface{}, topic string, cfg ConsumerConfig) {
	method := reflect.ValueOf(consumer).MethodByName("Process")

	r.eventProcessors[topic] = method
	r.eventConsumers[topic] = NewConsumer(cfg, topic)
}

func (r *ConsumerRegistry) Process(ctx context.Context) error {
	for topic, consumer := range r.eventConsumers {
		msg, err := consumer.Consume(context.Background()) // nolint:contextcheck
		if err != nil {
			log.Fatal(err)
		}

		messageType := r.eventProcessors[topic].Type().In(1)
		inArgPtr := reflect.New(messageType.Elem())

		in := inArgPtr.Interface()

		if err := protojson.Unmarshal(msg, in.(proto.Message)); err != nil {
			return err
		}

		res := r.eventProcessors[topic].Call([]reflect.Value{reflect.ValueOf(ctx), inArgPtr})

		if v := res[0].Interface(); v != nil {
			if err, ok := v.(error); ok {
				log.Fatal(err)
			}
		}
	}

	return nil
}
