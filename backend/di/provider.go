package di

import (
	"github.com/fabl3ss/banking_system/config"
	"github.com/fabl3ss/banking_system/modules/account"
	"github.com/fabl3ss/banking_system/modules/customer"
	"github.com/fabl3ss/banking_system/modules/manager"
	"github.com/fabl3ss/banking_system/modules/transfer"
	"github.com/fabl3ss/banking_system/pkg/consumer"
	"github.com/fabl3ss/banking_system/pkg/db"
	"github.com/fabl3ss/banking_system/pkg/producer"
	"github.com/fabl3ss/banking_system/pkg/redis"
	"github.com/fabl3ss/banking_system/projection_builder"
	"go.uber.org/fx"
)

func ProvideModules() []fx.Option {
	modules := []fx.Option{
		config.Module,
		db.Module,
		customer.Module,
		producer.Module,
		projection_builder.Module,
		redis.Module,
		account.Module,
		manager.Module,
		transfer.Module,
		consumer.Module,
	}

	return modules
}
