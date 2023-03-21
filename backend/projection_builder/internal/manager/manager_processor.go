package customer

import (
	"context"
	"database/sql"
	"fmt"

	customer_events "github.com/fabl3ss/banking_system/api/customer/events"
	"github.com/fabl3ss/banking_system/pkg/consumer"
)

type ManagerProcessor struct {
	db        *sql.DB
	tableName string
}

func NewManagerProcessor(db *sql.DB) *ManagerProcessor {
	return &ManagerProcessor{
		db:        db,
		tableName: "managers",
	}
}

func RegisterCustomerConsumer(
	consumer *ManagerProcessor,
	registry *consumer.ConsumerRegistry,
	cfg consumer.ConsumerConfig,
) {
	registry.Register(consumer, "customer-events", cfg)
}

func (c *ManagerProcessor) Process(ctx context.Context, message *customer_events.CustomerEvent) error {
	switch e := message.Event.(type) {
	case *customer_events.CustomerEvent_CustomerCreated:
		_, err := c.db.ExecContext(ctx,
			fmt.Sprintf("INSERT INTO %s (id, email, password_hash, created_at) VALUES (?, ?, ?, ?)", c.tableName),
			message.CustomerId,
			e.CustomerCreated.Email,
			e.CustomerCreated.PasswordHash,
			e.CustomerCreated.CreatedAt,
		)
		if err != nil {
			return err
		}

	case *customer_events.CustomerEvent_CustomerUpdated:
		_, err := c.db.ExecContext(ctx,
			fmt.Sprintf("UPDATE %s SET email = ?, password_hash = ?, created_at = ? WHERE id = ?", c.tableName),
			e.CustomerUpdated.Email,
			e.CustomerUpdated.PasswordHash,
			e.CustomerUpdated.UpdatedAt,
			message.CustomerId,
		)
		if err != nil {
			return err
		}
	}

	return nil
}
