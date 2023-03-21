package repository

import (
	"context"

	account_events "github.com/fabl3ss/banking_system/api/account/events"
	"github.com/fabl3ss/banking_system/api/common"
	account_model "github.com/fabl3ss/banking_system/modules/account/internal/model"
	"github.com/fabl3ss/banking_system/pkg/producer"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AccountRepository struct {
	eventProducer *producer.Producer
}

func NewAccountRepository(eventProducer *producer.Producer) *AccountRepository {
	return &AccountRepository{
		eventProducer: eventProducer,
	}
}

func (r *AccountRepository) Create(ctx context.Context, account *account_model.Account) error {
	event := &account_events.AccountEvent{
		AccountId: account.Id().String(),
		Event: &account_events.AccountEvent_AccountCreated{
			AccountCreated: &account_events.AccountCreated{
				HolderId: account.HolderId().String(),
				Balance: &common.Money{
					CurrencyCode: account.Balance().CurrencyCode,
					Amount:       account.Balance().Amount,
				},
				OpenedAt:   timestamppb.New(account.OpenedAt()),
				ExpiryDate: timestamppb.New(account.ExpiryDate()),
			},
		},
	}

	return r.eventProducer.Publish(ctx, event, "account-events", account.Id().String())
}

func (r *AccountRepository) Close(ctx context.Context, accountId uuid.UUID) error {
	event := &account_events.AccountEvent{
		AccountId: accountId.String(),
		Event: &account_events.AccountEvent_AccountClosed{
			AccountClosed: &account_events.AccountClosed{
				ClosedTime: timestamppb.Now(),
			},
		},
	}

	return r.eventProducer.Publish(ctx, event, "account-events", accountId.String())
}
