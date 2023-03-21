package domain

import (
	"context"
	"errors"

	customer_model "github.com/fabl3ss/banking_system/modules/customer/internal/model"
	"github.com/fabl3ss/banking_system/pkg/bcrypt"
)

type AuthenticationService struct {
	customerProjectionDAO customerProjectionDAO
}

func NewAuthenticationService(customerProjectionDAO customerProjectionDAO) *AuthenticationService {
	return &AuthenticationService{
		customerProjectionDAO: customerProjectionDAO,
	}
}

func (s *AuthenticationService) Authenticate(ctx context.Context, email string, password string) error {
	customer, err := s.customerProjectionDAO.GetByEmail(ctx, email)
	if err != nil {
		return err
	}

	if !bcrypt.IsPasswordEqualToHash(password, customer.PasswordHash()) {
		return errors.New("invalid password")
	}

	return nil
}

type customerProjectionDAO interface {
	GetByEmail(ctx context.Context, email string) (*customer_model.Customer, error)
}
