package manager

import (
	manager_repository "github.com/fabl3ss/banking_system/modules/manager/internal/persistence/repository"

	"go.uber.org/fx"
)

var Module = fx.Module("manager",
	fx.Provide(
		manager_repository.NewManagerRepository,
	),
)
