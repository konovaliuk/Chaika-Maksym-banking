package repository

import (
	"context"

	"github.com/fabl3ss/banking_system/api/common"
	transfer_events "github.com/fabl3ss/banking_system/api/transfer/events"
	transaction_model "github.com/fabl3ss/banking_system/modules/transfer/internal/model"
	"github.com/fabl3ss/banking_system/pkg/producer"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TransferRepository struct {
	eventProducer *producer.Producer
}

func NewTransferRepository(
	eventProducer *producer.Producer,
) *TransferRepository {
	return &TransferRepository{
		eventProducer: eventProducer,
	}
}

func (r *TransferRepository) Create(ctx context.Context, transfer *transaction_model.Transfer) error {
	event := &transfer_events.TransferEvent{
		TransferId: transfer.Id().String(),
		Event: &transfer_events.TransferEvent_TransferCreated{
			TransferCreated: &transfer_events.TransferCreated{
				SenderAccountId:    transfer.SenderAccountId().String(),
				RecipientAccountId: transfer.RecipientAccountId().String(),
				Funds: &common.Money{
					CurrencyCode: transfer.CurrencyCode(),
					Amount:       transfer.Amount(),
				},
				CreatedAt: timestamppb.New(transfer.CreatedAt()),
				Status:    mapTransferStatusToDto(transfer.Status()),
			},
		},
	}

	return r.eventProducer.Publish(ctx, event, "funds-transfer-events", transfer.Id().String())
}

func mapTransferStatusToDto(status transaction_model.TransferStatus) transfer_events.TransferStatus {
	switch status {
	case transaction_model.TransferStatusPending:
		return transfer_events.TransferStatus_TRANSFER_STATUS_PENDING
	case transaction_model.TransferStatusApproved:
		return transfer_events.TransferStatus_TRANSFER_STATUS_APPROVED
	case transaction_model.TransferStatusDeclined:
		return transfer_events.TransferStatus_TRANSFER_STATUS_DECLINED
	}

	return transfer_events.TransferStatus_TRANSFER_STATUS_UNDEFINED
}
