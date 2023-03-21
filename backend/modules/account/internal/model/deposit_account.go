package model

type DepositAccount struct {
	Account
	depositAmount int64
	annualRate    int
}

func NewDepositAccount(
	account Account,
	depositAmount int64,
	annualRate int,
) *DepositAccount {
	return &DepositAccount{
		Account:       account,
		depositAmount: depositAmount,
		annualRate:    annualRate,
	}
}

func (a *DepositAccount) DepositAmount() int64 {
	return a.depositAmount
}

func (a *DepositAccount) SetDepositAmount(depositAmount int64) {
	a.depositAmount = depositAmount
}

func (a *DepositAccount) AnnualRate() int {
	return a.annualRate
}

func (a *DepositAccount) SetAnnualRate(annualRate int) {
	a.annualRate = annualRate
}
