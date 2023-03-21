package projection_builder

import (
	"github.com/fabl3ss/banking_system/pkg/consumer"
	"github.com/fabl3ss/banking_system/projection_builder/internal/account"
	"github.com/fabl3ss/banking_system/projection_builder/internal/customer"
	"github.com/fabl3ss/banking_system/projection_builder/internal/transfer"

	"go.uber.org/fx"
)

var Module = fx.Module("projection_builder",
	fx.Provide(
		customer.NewCustomerProcessor,
		account.NewAccountProcessor,
		transfer.NewTransferProcessor,
	),
	fx.Invoke(func(
		cfg consumer.ConsumerConfig,
		registry *consumer.ConsumerRegistry,
		customerConsumer *customer.CustomerProcessor,
		accountConsumer *account.AccountProcessor,
		transferConsumer *transfer.TransferProcessor,
	) {
		customer.RegisterCustomerConsumer(customerConsumer, registry, cfg)
	}),
)
