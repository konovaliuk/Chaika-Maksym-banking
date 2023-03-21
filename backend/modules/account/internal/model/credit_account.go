package model

type CreditAccount struct {
	Account

	limit           int64
	debt            int64
	accruedInterest int
	creditRate      int
}

func NewCreditAccount(
	account Account,
	limit int64,
	debt int64,
	accruedInterest int,
	creditRate int,
) *CreditAccount {
	return &CreditAccount{
		Account:         account,
		limit:           limit,
		debt:            debt,
		accruedInterest: accruedInterest,
		creditRate:      creditRate,
	}
}

func (a *CreditAccount) Limit() int64 {
	return a.limit
}

func (a *CreditAccount) SetLimit(limit int64) {
	a.limit = limit
}

func (a *CreditAccount) Debt() int64 {
	return a.debt
}

func (a *CreditAccount) SetDebt(debt int64) {
	a.debt = debt
}

func (a *CreditAccount) AccruedInterest() int {
	return a.accruedInterest
}

func (a *CreditAccount) SetAccruedInterest(accruedInterest int) {
	a.accruedInterest = accruedInterest
}

func (a *CreditAccount) CreditRate() int {
	return a.creditRate
}

func (a *CreditAccount) SetCreditRate(creditRate int) {
	a.creditRate = creditRate
}
