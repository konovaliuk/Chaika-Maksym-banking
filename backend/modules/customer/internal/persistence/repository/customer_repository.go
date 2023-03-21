package repository

import (
	"context"

	customer_events "github.com/fabl3ss/banking_system/api/customer/events"
	customer_model "github.com/fabl3ss/banking_system/modules/customer/internal/model"

	"github.com/fabl3ss/banking_system/pkg/producer"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CustomerRepository struct {
	cacheVerifier CustomerCacheVerifierDAO
	eventProducer *producer.Producer
}

func NewCustomerRepository(
	cacheVerifier CustomerCacheVerifierDAO,
	eventProducer *producer.Producer,
) *CustomerRepository {
	return &CustomerRepository{
		cacheVerifier: cacheVerifier,
		eventProducer: eventProducer,
	}
}

func (r *CustomerRepository) Create(ctx context.Context, customer *customer_model.Customer) error {
	if err := r.cacheVerifier.Insert(ctx, customer); err != nil {
		return err
	}

	event := &customer_events.CustomerEvent{
		CustomerId: customer.Id().String(),
		Event: &customer_events.CustomerEvent_CustomerCreated{
			CustomerCreated: &customer_events.CustomerCreated{
				Email:        customer.Email(),
				PasswordHash: customer.PasswordHash(),
				CreatedAt:    timestamppb.New(customer.CreatedAt()),
			},
		},
	}

	return r.eventProducer.Publish(ctx, event, "customer-events", customer.Id().String())
}

func (r *CustomerRepository) Update(ctx context.Context, customer *customer_model.Customer) error {
	if err := r.cacheVerifier.Insert(ctx, customer); err != nil {
		return err
	}

	event := &customer_events.CustomerEvent{
		CustomerId: customer.Id().String(),
		Event: &customer_events.CustomerEvent_CustomerUpdated{
			CustomerUpdated: &customer_events.CustomerUpdated{
				Email:        customer.Email(),
				PasswordHash: customer.PasswordHash(),
				UpdatedAt:    timestamppb.New(customer.UpdatedAt()),
			},
		},
	}

	return r.eventProducer.Publish(ctx, event, "customer-events", customer.Id().String())
}

type CustomerCacheVerifierDAO interface {
	Insert(ctx context.Context, customer *customer_model.Customer) error
	Update(ctx context.Context, customer *customer_model.Customer) error
}
