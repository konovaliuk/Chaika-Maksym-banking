package account

import (
	account_dao "github.com/fabl3ss/banking_system/modules/account/internal/persistence/dao"
	account_repository "github.com/fabl3ss/banking_system/modules/account/internal/persistence/repository"
	"go.uber.org/fx"
)

var Module = fx.Module("account",
	fx.Provide(
		account_repository.NewAccountRepository,

		account_dao.NewAccountProjectionDAO,
	),
)
