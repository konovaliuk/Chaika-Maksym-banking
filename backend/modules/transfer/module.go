package transfer

import (
	transfer_dao "github.com/fabl3ss/banking_system/modules/transfer/internal/persistence/dao"
	transaction_repository "github.com/fabl3ss/banking_system/modules/transfer/internal/persistence/repository"

	"go.uber.org/fx"
)

var Module = fx.Module("transfer",
	fx.Provide(
		transaction_repository.NewTransferRepository,
		transfer_dao.NewTransferProjectionDAO,
	),
)
