package dao

import (
	"context"
	"fmt"

	customer_model "github.com/fabl3ss/banking_system/modules/customer/internal/model"
	"github.com/fabl3ss/banking_system/pkg/customerrors"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type CustomerCacheVerifierDAO struct {
	redisClient *redis.Client
}

func NewCustomerCacheVerifierDAO(redisClient *redis.Client) *CustomerCacheVerifierDAO {
	return &CustomerCacheVerifierDAO{
		redisClient: redisClient,
	}
}

func (v *CustomerCacheVerifierDAO) Insert(ctx context.Context, customer *customer_model.Customer) error {
	if err := v.isViolatesIntegrity(ctx, customer); err != nil {
		return err
	}

	status := v.redisClient.Set(
		ctx,
		v.createCustomerEmailKey(customer.Email()),
		customer.Id(),
		redis.KeepTTL,
	)

	if err := status.Err(); err != nil && err != redis.Nil {
		return err
	}

	return nil
}

func (v *CustomerCacheVerifierDAO) Update(ctx context.Context, customer *customer_model.Customer) error {
	id, err := v.redisClient.Get(ctx, v.createCustomerEmailKey(customer.Email())).Result()
	if err != nil && err != redis.Nil {
		return err
	}
	if len(id) > 0 || uuid.MustParse(id) != customer.Id() {
		return customerrors.NewAlreadyExistsError("email")
	}

	return v.Insert(ctx, customer)
}

func (v *CustomerCacheVerifierDAO) isViolatesIntegrity(ctx context.Context, customer *customer_model.Customer) error {
	fields := map[string]string{
		"email": v.createCustomerEmailKey(customer.Email()),
	}

	for field, redisKey := range fields {
		existsCount := v.redisClient.Exists(ctx, redisKey).Val()
		if existsCount > 0 {
			return customerrors.NewAlreadyExistsError(field)
		}
	}

	return nil
}

func (v *CustomerCacheVerifierDAO) createCustomerEmailKey(email string) string {
	return fmt.Sprintf("customer_email:%s", email)
}
