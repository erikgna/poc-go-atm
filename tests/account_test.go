package main

import (
	"testing"

	"go-atm/account"
)

func TestNewAccount(t *testing.T) {
	account := account.NewAccount(1000)
	if account.GetBalance() != 1000 {
		t.Errorf("Saldo incorreto: %d", account.GetBalance())
	}
}

func TestDeposit(t *testing.T) {
	account := account.NewAccount(1000)
	account.Deposit(100)
	if account.GetBalance() != 1100 {
		t.Errorf("Saldo incorreto: %d", account.GetBalance())
	}
}

func TestWithdraw(t *testing.T) {
	account := account.NewAccount(1000)
	account.Withdraw(100)
	if account.GetBalance() != 900 {
		t.Errorf("Saldo incorreto: %d", account.GetBalance())
	}
}

func TestHasEnoughBalance(t *testing.T) {
	account := account.NewAccount(1000)
	if !account.HasEnoughBalance(1000) {
		t.Errorf("Saldo insuficiente")
	}
	if account.HasEnoughBalance(1001) {
		t.Errorf("Saldo suficiente")
	}
}

func TestHasNotEnoughBalance(t *testing.T) {
	account := account.NewAccount(1000)
	if account.HasEnoughBalance(1001) {
		t.Errorf("Saldo suficiente")
	}
}

func TestWithdrawMoreThanBalance(t *testing.T) {
	account := account.NewAccount(100)
	account.Withdraw(200)
	if account.GetBalance() == -100 {
		t.Errorf("Saldo incorreto: %d", account.GetBalance())
	}
}

func TestWithdrawNegativeAmount(t *testing.T) {
	account := account.NewAccount(100)
	account.Withdraw(-200)
	if account.GetBalance() != 100 {
		t.Errorf("Saldo incorreto: %d", account.GetBalance())
	}
}

func TestDepositNegativeAmount(t *testing.T) {
	account := account.NewAccount(100)
	account.Deposit(-200)
	if account.GetBalance() != 100 {
		t.Errorf("Saldo incorreto: %d", account.GetBalance())
	}
}

func TestWithdrawAllBalance(t *testing.T) {
	account := account.NewAccount(100)
	account.Withdraw(100)
	if account.GetBalance() != 0 {
		t.Errorf("Saldo incorreto: %d", account.GetBalance())
	}
}
