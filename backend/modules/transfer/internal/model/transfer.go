package model

import (
	"time"

	"github.com/google/uuid"
)

type TransferStatus string

const (
	TransferStatusPending  TransferStatus = "pending"
	TransferStatusApproved TransferStatus = "approved"
	TransferStatusDeclined TransferStatus = "declined"
)

type Transfer struct {
	id                 uuid.UUID
	senderAccountId    uuid.UUID
	recipientAccountId uuid.UUID
	amount             int64
	currencyCode       string
	createdAt          time.Time
	updatedAt          time.Time
	status             TransferStatus
}

func NewTransfer(
	id uuid.UUID,
	senderAccountId uuid.UUID,
	recipientAccountId uuid.UUID,
	amount int64,
	currencyCode string,
	createdAt time.Time,
	updatedAt time.Time,
	status TransferStatus,
) *Transfer {
	return &Transfer{
		id:                 id,
		senderAccountId:    senderAccountId,
		recipientAccountId: recipientAccountId,
		amount:             amount,
		currencyCode:       currencyCode,
		createdAt:          createdAt,
		updatedAt:          updatedAt,
		status:             status,
	}
}

func (t *Transfer) Id() uuid.UUID {
	return t.id
}

func (t *Transfer) SetId(id uuid.UUID) {
	t.id = id
}

func (t *Transfer) SenderAccountId() uuid.UUID {
	return t.senderAccountId
}

func (t *Transfer) SetSenderAccountId(senderAccountId uuid.UUID) {
	t.senderAccountId = senderAccountId
}

func (t *Transfer) RecipientAccountId() uuid.UUID {
	return t.recipientAccountId
}

func (t *Transfer) SetRecipientAccountId(recipientAccountId uuid.UUID) {
	t.recipientAccountId = recipientAccountId
}

func (t *Transfer) Amount() int64 {
	return t.amount
}

func (t *Transfer) SetAmount(amount int64) {
	t.amount = amount
}

func (t *Transfer) CurrencyCode() string {
	return t.currencyCode
}

func (t *Transfer) SetCurrencyCode(currencyCode string) {
	t.currencyCode = currencyCode
}

func (t *Transfer) CreatedAt() time.Time {
	return t.createdAt
}

func (t *Transfer) SetCreatedAt(createdAt time.Time) {
	t.createdAt = createdAt
}

func (t *Transfer) UpdatedAt() time.Time {
	return t.createdAt
}

func (t *Transfer) SetUpdatedAt(createdAt time.Time) {
	t.createdAt = createdAt
}

func (t *Transfer) Status() TransferStatus {
	return t.status
}

func (t *Transfer) SetStatus(status TransferStatus) {
	t.status = status
}
