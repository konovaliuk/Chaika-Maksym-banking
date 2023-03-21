package transfer

import (
	"context"
	"database/sql"
	"fmt"

	transfer_events "github.com/fabl3ss/banking_system/api/transfer/events"
	"github.com/fabl3ss/banking_system/pkg/consumer"
	"github.com/fabl3ss/banking_system/pkg/producer"
	"github.com/google/uuid"
)

type TransferProcessor struct {
	db            *sql.DB
	tableName     string
	eventProducer *producer.Producer
}

func NewTransferProcessor(db *sql.DB, eventProducer *producer.Producer) *TransferProcessor {
	return &TransferProcessor{
		db:            db,
		tableName:     "transfers",
		eventProducer: eventProducer,
	}
}

func RegisterTransferConsumer(
	consumer *TransferProcessor,
	registry *consumer.ConsumerRegistry,
	cfg consumer.ConsumerConfig,
) {
	registry.Register(consumer, "funds-transfer-events", cfg)
}

func (p *TransferProcessor) Process(ctx context.Context, message *transfer_events.TransferEvent) error {
	switch e := message.Event.(type) {
	case *transfer_events.TransferEvent_TransferCreated:
		_, err := p.db.ExecContext(ctx,
			fmt.Sprintf(`INSERT INTO %s (
                       id, 
                       sender_account_id, 
                       recipient_account_id, 
                       amount, 
                       currency_code, 
                       created_at, 
                       updated_at, 
                       status
                  ) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`, p.tableName),
			uuid.MustParse(message.TransferId),
			uuid.MustParse(e.TransferCreated.SenderAccountId),
			uuid.MustParse(e.TransferCreated.RecipientAccountId),
			e.TransferCreated.Funds.Amount,
			e.TransferCreated.Funds.CurrencyCode,
			e.TransferCreated.CreatedAt.AsTime(),
			e.TransferCreated.CreatedAt.AsTime(),
		)
		if err != nil {
			event := &transfer_events.TransferEvent{
				TransferId: message.TransferId,
				Event: &transfer_events.TransferEvent_TransferCreateFailed{
					TransferCreateFailed: &transfer_events.TransferCreateFailed{
						Description: err.Error(),
					},
				},
			}

			return p.eventProducer.Publish(ctx, event, "funds-transfer-events", message.TransferId)
		}

	case *transfer_events.TransferEvent_TransferStatusUpdated:
		_, err := p.db.ExecContext(ctx,
			fmt.Sprintf("UPDATE %s SET status = ?, updated_at = ? WHERE id = ?", p.tableName),
			e.TransferStatusUpdated.Status,
			e.TransferStatusUpdated.UpdatedAt.AsTime(),
			uuid.MustParse(message.TransferId),
		)
		if err != nil {
			event := &transfer_events.TransferEvent{
				TransferId: message.TransferId,
				Event: &transfer_events.TransferEvent_TransferStatusUpdateFiled{
					TransferStatusUpdateFiled: &transfer_events.TransferStatusUpdateFailed{
						Description: err.Error(),
					},
				},
			}

			return p.eventProducer.Publish(ctx, event, "funds-transfer-events", message.TransferId)
		}
	}

	return nil
}
