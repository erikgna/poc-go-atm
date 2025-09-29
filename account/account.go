package account

import "errors"

type Account struct {
	Balance int
}

func NewAccount(initialBalance int) *Account {
	return &Account{Balance: initialBalance}
}

func (a *Account) Deposit(amount int) error {
	if amount < 0 {
		return errors.New("amount is negative")
	}
	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount int) error {
	if amount < 0 {
		return errors.New("amount is negative")
	}
	if amount > a.Balance {
		return errors.New("insufficient balance")
	}
	a.Balance -= amount
	return nil
}

func (a *Account) GetBalance() int {
	return a.Balance
}

func (a *Account) HasEnoughBalance(amount int) bool {
	return a.Balance >= amount
}
