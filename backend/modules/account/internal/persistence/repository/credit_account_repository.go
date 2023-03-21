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

type CreditAccountRepository struct {
	eventProducer *producer.Producer
}

func NewCreditAccountRepository(eventProducer *producer.Producer) *CreditAccountRepository {
	return &CreditAccountRepository{
		eventProducer: eventProducer,
	}
}

func (r *CreditAccountRepository) Create(ctx context.Context, account *account_model.CreditAccount) error {
	event := &account_events.CreditAccountEvent{
		AccountId: account.Id().String(),
		Event: &account_events.CreditAccountEvent_AccountCreated{
			AccountCreated: &account_events.CreditAccountCreated{
				HolderId: account.HolderId().String(),
				Balance: &common.Money{
					CurrencyCode: account.Balance().CurrencyCode,
					Amount:       account.Balance().Amount,
				},
				OpenedAt:        timestamppb.New(account.OpenedAt()),
				ExpiryDate:      timestamppb.New(account.ExpiryDate()),
				Limit:           account.Limit(),
				Debt:            account.Debt(),
				AccruedInterest: int32(account.AccruedInterest()),
				CreditRate:      int32(account.CreditRate()),
			},
		},
	}

	return r.eventProducer.Publish(ctx, event, "account-events", account.Id().String())
}

func (r *CreditAccountRepository) Close(ctx context.Context, accountId uuid.UUID) error {
	event := &account_events.CreditAccountEvent{
		AccountId: accountId.String(),
		Event: &account_events.CreditAccountEvent_AccountClosed{
			AccountClosed: &account_events.CreditAccountClosed{
				ClosedTime: timestamppb.Now(),
			},
		},
	}

	return r.eventProducer.Publish(ctx, event, "account-events", accountId.String())
}
