package config

import (
	"github.com/caarlos0/env/v7"
	"github.com/fabl3ss/banking_system/pkg/consumer"
	"github.com/fabl3ss/banking_system/pkg/db"
	"github.com/fabl3ss/banking_system/pkg/producer"
	"github.com/fabl3ss/banking_system/pkg/redis"
)

type Config struct {
	DB       db.Config
	Producer producer.ProducerConfig
	Consumer consumer.ConsumerConfig
	Redis    redis.Config
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
