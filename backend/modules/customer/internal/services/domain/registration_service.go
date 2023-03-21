package domain

import (
	"context"
	"time"

	customer_model "github.com/fabl3ss/banking_system/modules/customer/internal/model"
	"github.com/fabl3ss/banking_system/pkg/bcrypt"
	"github.com/google/uuid"
)

type RegistrationService struct {
	customerRepository customerRepository
}

func NewRegistrationService(customerRepository customerRepository) *RegistrationService {
	return &RegistrationService{
		customerRepository: customerRepository,
	}
}

func (s *RegistrationService) Register(ctx context.Context, email string, password string) error {
	passwordHash, err := bcrypt.HashPassword(password)
	if err != nil {
		return err
	}

	customerModel := customer_model.NewCustomer(
		uuid.New(),
		email,
		passwordHash,
		time.Now(),
		time.Now(),
	)

	return s.customerRepository.Create(ctx, customerModel)
}

type customerRepository interface {
	Create(ctx context.Context, customer *customer_model.Customer) error
	Update(ctx context.Context, customer *customer_model.Customer) error
}
