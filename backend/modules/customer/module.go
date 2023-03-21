package customer

import (
	"github.com/fabl3ss/banking_system/modules/customer/customer_tests"
	customer_dao "github.com/fabl3ss/banking_system/modules/customer/internal/persistence/dao"
	customer_repository "github.com/fabl3ss/banking_system/modules/customer/internal/persistence/repository"
	"github.com/fabl3ss/banking_system/modules/customer/internal/services/domain"
	"go.uber.org/fx"
)

var Module = fx.Module("customer",
	fx.Provide(
		customer_dao.NewCustomerProjectionDAO,
		customer_dao.NewCustomerCacheVerifierDAO,

		customer_repository.NewCustomerRepository,

		domain.NewRegistrationService,
		domain.NewAuthenticationService,

		customer_tests.NewCustomerStepHandler,
	),
)
