package account

import (
	"context"
	"database/sql"
	"fmt"

	account_events "github.com/fabl3ss/banking_system/api/account/events"
	"github.com/fabl3ss/banking_system/pkg/consumer"
	"github.com/fabl3ss/banking_system/pkg/producer"
	"github.com/google/uuid"
)

type AccountProcessor struct {
	db            *sql.DB
	tableName     string
	eventProducer *producer.Producer
}

func NewAccountProcessor(db *sql.DB, eventProducer *producer.Producer) *AccountProcessor {
	return &AccountProcessor{
		db:            db,
		tableName:     "accounts",
		eventProducer: eventProducer,
	}
}

func RegisterAccountConsumer(
	consumer *AccountProcessor,
	registry *consumer.ConsumerRegistry,
	cfg consumer.ConsumerConfig,
) {
	registry.Register(consumer, "account-events", cfg)
}

func (p *AccountProcessor) Process(ctx context.Context, message *account_events.AccountEvent) error {
	switch e := message.Event.(type) {
	case *account_events.AccountEvent_AccountCreated:
		_, err := p.db.ExecContext(ctx,
			fmt.Sprintf(`INSERT INTO %s (
                      id, 
                      holder_id, 
                      currency_code, 
                      amount,
                      opened_at, 
                      expiry_data
                   ) VALUES (?, ?, ?, ?)`, p.tableName),
			uuid.MustParse(message.AccountId),
			uuid.MustParse(e.AccountCreated.HolderId),
			e.AccountCreated.Balance.CurrencyCode,
			e.AccountCreated.Balance.Amount,
			e.AccountCreated.OpenedAt.AsTime(),
			e.AccountCreated.ExpiryDate.AsTime(),
		)
		if err != nil {
			event := &account_events.AccountEvent{
				AccountId: message.AccountId,
				Event: &account_events.AccountEvent_AccountCreateFailed{
					AccountCreateFailed: &account_events.AccountCreateFailed{
						Description: err.Error(),
					},
				},
			}

			return p.eventProducer.Publish(ctx, event, "account-events", message.AccountId)
		}

	case *account_events.AccountEvent_AccountApproved:
		_, err := p.db.ExecContext(ctx,
			fmt.Sprintf("UPDATE %s SET status = 'approved', updated_at = ? WHERE id = ?", p.tableName),
			e.AccountApproved.ApprovalTime.AsTime(),
			message.AccountId,
		)
		if err != nil {
			event := &account_events.AccountEvent{
				AccountId: message.AccountId,
				Event: &account_events.AccountEvent_AccountApproveFailed{
					AccountApproveFailed: &account_events.AccountApproveFailed{
						Description: err.Error(),
					},
				},
			}

			return p.eventProducer.Publish(ctx, event, "account-events", message.AccountId)
		}

	case *account_events.AccountEvent_AccountClosed:
		_, err := p.db.ExecContext(ctx,
			fmt.Sprintf("UPDATE %s SET status = 'closed', updated_at = ? WHERE id = ?", p.tableName),
			e.AccountClosed.ClosedTime.AsTime(),
			message.AccountId,
		)
		if err != nil {
			event := &account_events.AccountEvent{
				AccountId: message.AccountId,
				Event: &account_events.AccountEvent_AccountCloseFailed{
					AccountCloseFailed: &account_events.AccountCloseFailed{
						Description: err.Error(),
					},
				},
			}

			return p.eventProducer.Publish(ctx, event, "account-events", message.AccountId)
		}
	}

	return nil
}
