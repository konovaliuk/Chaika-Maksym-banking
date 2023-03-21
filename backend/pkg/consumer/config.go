package consumer

type ConsumerConfig struct {
	Addr string `env:"KAFKA_ADDR" envDefault:"localhost:9093"`
}
