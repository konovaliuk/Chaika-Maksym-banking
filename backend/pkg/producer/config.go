package producer

type ProducerConfig struct {
	Addr string `env:"KAFKA_ADDR" envDefault:"localhost:9093"`
}
