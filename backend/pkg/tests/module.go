package tests

import (
	"github.com/fabl3ss/banking_system/pkg/tests/step_handlers"
	"go.uber.org/fx"
)

var Module = fx.Module("customer_tests",
	fx.Provide(step_handlers.NewDBStepHandler),
)
