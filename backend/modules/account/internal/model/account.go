package model

import (
	"time"

	"github.com/google/uuid"
)

type AccountType string

const (
	AccountTypeCredit  = "credit"
	AccountTypeDeposit = "deposit"
)

type Account struct {
	id         uuid.UUID
	holderId   uuid.UUID
	balance    Money
	openedAt   time.Time
	expiryDate time.Time
}

func NewAccount(
	id uuid.UUID,
	holderId uuid.UUID,
	currencyCode string,
	amount int64,
	openedAt time.Time,
	expiryDate time.Time,
) *Account {
	return &Account{
		id:       id,
		holderId: holderId,
		balance: Money{
			CurrencyCode: currencyCode,
			Amount:       amount,
		},
		openedAt:   openedAt,
		expiryDate: expiryDate,
	}
}

func (a *Account) Id() uuid.UUID {
	return a.id
}

func (a *Account) SetId(id uuid.UUID) {
	a.id = id
}

func (a *Account) HolderId() uuid.UUID {
	return a.holderId
}

func (a *Account) SetHolderId(holderId uuid.UUID) {
	a.holderId = holderId
}

func (a *Account) Balance() Money {
	return a.balance
}

func (a *Account) SetBalance(balance Money) {
	a.balance = balance
}

func (a *Account) OpenedAt() time.Time {
	return a.openedAt
}

func (a *Account) SetOpenedAt(openedAt time.Time) {
	a.openedAt = openedAt
}

func (a *Account) ExpiryDate() time.Time {
	return a.expiryDate
}

func (a *Account) SetExpiryDate(expiryDate time.Time) {
	a.expiryDate = expiryDate
}
