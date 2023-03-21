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

type DepositAccountRepository struct {
	eventProducer *producer.Producer
}

func NewDepositAccountRepository(eventProducer *producer.Producer) *DepositAccountRepository {
	return &DepositAccountRepository{
		eventProducer: eventProducer,
	}
}

func (r *DepositAccountRepository) Create(ctx context.Context, account *account_model.DepositAccount) error {
	event := &account_events.DepositAccountEvent{
		AccountId: account.Id().String(),
		Event: &account_events.DepositAccountEvent_AccountCreated{
			AccountCreated: &account_events.DepositAccountCreated{
				HolderId: account.HolderId().String(),
				Balance: &common.Money{
					CurrencyCode: account.Balance().CurrencyCode,
					Amount:       account.Balance().Amount,
				},
				OpenedAt:      timestamppb.New(account.OpenedAt()),
				ExpiryDate:    timestamppb.New(account.ExpiryDate()),
				DepositAmount: account.DepositAmount(),
				AnnualRate:    int32(account.AnnualRate()),
			},
		},
	}

	return r.eventProducer.Publish(ctx, event, "account-events", account.Id().String())
}

func (r *DepositAccountRepository) Close(ctx context.Context, accountId uuid.UUID) error {
	event := &account_events.DepositAccountEvent{
		AccountId: accountId.String(),
		Event: &account_events.DepositAccountEvent_AccountClosed{
			AccountClosed: &account_events.DepositAccountClosed{
				ClosedTime: timestamppb.Now(),
			},
		},
	}

	return r.eventProducer.Publish(ctx, event, "account-events", accountId.String())
}
