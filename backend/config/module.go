package config

import (
	"github.com/fabl3ss/banking_system/pkg/consumer"
	"github.com/fabl3ss/banking_system/pkg/db"
	"github.com/fabl3ss/banking_system/pkg/producer"
	"github.com/fabl3ss/banking_system/pkg/redis"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	NewConfig,
	func(cfg *Config) db.Config {
		return cfg.DB
	},
	func(cfg *Config) producer.ProducerConfig {
		return cfg.Producer
	},
	func(cfg *Config) consumer.ConsumerConfig {
		return cfg.Consumer
	},
	func(cfg *Config) redis.Config {
		return cfg.Redis
	},
)
