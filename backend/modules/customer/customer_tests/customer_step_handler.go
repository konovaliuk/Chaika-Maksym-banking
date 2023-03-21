package customer_tests

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"sync"
	"time"

	"github.com/cucumber/godog"
	customer_model "github.com/fabl3ss/banking_system/modules/customer/internal/model"
	customer_dao "github.com/fabl3ss/banking_system/modules/customer/internal/persistence/dao"
	customer_repository "github.com/fabl3ss/banking_system/modules/customer/internal/persistence/repository"
	"github.com/google/uuid"
)

type CustomerStepHandler struct {
	customerRepository    *customer_repository.CustomerRepository
	customerProjectionDAO *customer_dao.CustomerProjectionDAO
	methodResult          map[string]interface{}
	wg                    *sync.WaitGroup
}

func NewCustomerStepHandler(
	customerRepository *customer_repository.CustomerRepository,
	customerProjectionDAO *customer_dao.CustomerProjectionDAO,
) *CustomerStepHandler {
	return &CustomerStepHandler{
		customerRepository:    customerRepository,
		customerProjectionDAO: customerProjectionDAO,
	}
}

func (h *CustomerStepHandler) RegisterSteps(sc *godog.ScenarioContext, wg *sync.WaitGroup) {
	sc.Step(`^I execute "([^"]*)" method in "([^"]*)" with body:$`, h.executeMethod)
	sc.Step(`^customer method should return:$`, h.methodShouldReturn)
	h.wg = wg
}

func (h *CustomerStepHandler) executeMethod(methodName string, callerName string, body string) error {
	switch callerName {
	case "CustomerRepository":
		return h.executeRepositoryMethod(methodName, body)
	case "CustomerProjectionDAO":
		var param interface{}
		if err := json.Unmarshal([]byte(body), &param); err != nil {
			return err
		}

		return h.executeDaoMethod(methodName, param)
	}

	return fmt.Errorf("invalid caller name %s", callerName)
}

func (h *CustomerStepHandler) executeRepositoryMethod(methodName string, body string) error {
	var customerDto struct {
		Id           string
		Email        string
		PasswordHash string
		CreatedAt    string
		UpdatedAt    string
	}

	if err := json.Unmarshal([]byte(body), &customerDto); err != nil {
		return err
	}

	createdAt, err := time.Parse(time.RFC3339, customerDto.CreatedAt)
	if err != nil {
		return err
	}

	updatedAt, err := time.Parse(time.RFC3339, customerDto.CreatedAt)
	if err != nil {
		return err
	}

	customer := customer_model.NewCustomer(
		uuid.MustParse(customerDto.Id),
		customerDto.Email,
		customerDto.PasswordHash,
		createdAt,
		updatedAt,
	)

	method := reflect.ValueOf(h.customerRepository).MethodByName(methodName)
	res := method.Call([]reflect.Value{reflect.ValueOf(context.Background()), reflect.ValueOf(customer)})

	if v := res[0].Interface(); v != nil {
		if err, ok := v.(error); ok {
			log.Fatal(err)
		}
	}

	h.wg.Wait()
	return nil
}

func (h *CustomerStepHandler) executeDaoMethod(methodName string, parameter interface{}) error {
	var res []reflect.Value

	switch p := parameter.(type) {
	case map[string]interface{}:
		for _, v := range p {
			res = reflect.ValueOf(h.customerProjectionDAO).MethodByName(methodName).Call(
				[]reflect.Value{reflect.ValueOf(context.Background()), reflect.ValueOf(v)},
			)
		}
	default:
		return fmt.Errorf("unsupported parameter type %T", parameter)
	}

	if v := res[1].Interface(); v != nil {
		if err, ok := v.(error); ok {
			return err
		}
	}

	customerPtr := res[0].Interface().(*customer_model.Customer)

	h.methodResult = map[string]interface{}{
		"id":           customerPtr.Id().String(),
		"email":        customerPtr.Email(),
		"passwordHash": customerPtr.PasswordHash(),
		"createdAt":    customerPtr.CreatedAt().Format(time.RFC3339),
		"updatedAt":    customerPtr.UpdatedAt().Format(time.RFC3339),
	}

	return nil
}

func (h *CustomerStepHandler) methodShouldReturn(body string) error {
	var bodyMap map[string]interface{}
	if err := json.Unmarshal([]byte(body), &bodyMap); err != nil {
		return err
	}

	if !reflect.DeepEqual(bodyMap, h.methodResult) {
		return fmt.Errorf("invalid method response, Actual: %s\nExpected: %s", h.methodResult, body)
	}

	return nil
}
