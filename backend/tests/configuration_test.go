package tests

import (
	"context"
	"errors"
	"flag"
	"log"
	"os"
	"path"
	"runtime"
	"sync"
	"testing"
	"time"

	"github.com/cucumber/godog"
	"github.com/fabl3ss/banking_system/di"
	"github.com/fabl3ss/banking_system/modules/customer/customer_tests"
	"github.com/fabl3ss/banking_system/pkg/consumer"
	"github.com/fabl3ss/banking_system/pkg/db"
	"github.com/fabl3ss/banking_system/pkg/tests"
	"github.com/fabl3ss/banking_system/pkg/tests/step_handlers"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}

func TestAllScenarios(test *testing.T) {
	fxOptions := di.ProvideModules()
	fxOptions = append(
		fxOptions,
		tests.Module,
		fx.Invoke(
			func(
				lc fx.Lifecycle,
				dbStepHandler *step_handlers.DBStepHandler,
				customerStepHandler *customer_tests.CustomerStepHandler,
				dbCleaner *db.Cleaner,
				consumerRegistry *consumer.ConsumerRegistry,
				redisClient *redis.Client,
			) {
				wg := sync.WaitGroup{}
				go func() {
					wg.Add(1) // nolint:staticcheck
					if err := consumerRegistry.Process(context.Background()); err != nil {
						log.Fatal(err)
					}
					wg.Done()
				}()

				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						suite := godog.TestSuite{
							Name: "main",
							ScenarioInitializer: func(sc *godog.ScenarioContext) {
								dbStepHandler.RegisterSteps(sc)
								customerStepHandler.RegisterSteps(sc, &wg)
								redisClient.FlushAll(context.Background())
								if err := dbCleaner.CleanDatabase(); err != nil {
									log.Fatal(err)
								}
							},
							Options: &godog.Options{
								Format:    "progress",
								Strict:    true,
								Paths:     []string{"tests/scenarios"},
								Randomize: time.Now().UTC().UnixNano(),
								Tags:      *flag.String("godog.tags", "", ""), //nolint:gocritic
							},
						}

						status := suite.Run()
						if status != 0 {
							return errors.New("func customer_tests failed")
						}

						return nil
					},
					OnStop: func(ctx context.Context) error {
						return nil
					},
				})
			},
		),
	)

	app := fx.New(fxOptions...)

	if err := app.Start(context.Background()); err != nil {
		test.Error(err)
	}
}
