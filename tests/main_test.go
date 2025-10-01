package main

import (
	"reflect"
	"testing"

	"go-atm/account"
	"go-atm/atm"
)

func TestWithdrawWithNotes_Success(t *testing.T) {
	acct := account.NewAccount(1000)

	notes, err := atm.WithdrawWithNotes(240, acct)
	if err != nil {
		t.Fatalf("error to withdraw: %v", err)
	}

	expected := []atm.Notes{
		{Note: 100, Quantity: 2},
		{Note: 20, Quantity: 2},
	}

	if !reflect.DeepEqual(notes, expected) {
		t.Errorf("expected %v, received %v", expected, notes)
	}

	if acct.GetBalance() != 760 {
		t.Errorf("balance should be 760: %d", acct.GetBalance())
	}
}

func TestWithdrawWithNotes_ExactDenomination(t *testing.T) {
	acct := account.NewAccount(500)

	notes, err := atm.WithdrawWithNotes(50, acct)
	if err != nil {
		t.Fatalf("error to withdraw: %v", err)
	}

	expected := []atm.Notes{
		{Note: 50, Quantity: 1},
	}

	if !reflect.DeepEqual(notes, expected) {
		t.Errorf("expected %v, received %v", expected, notes)
	}

	if acct.GetBalance() != 450 {
		t.Errorf("balance should be 450: %d", acct.GetBalance())
	}
}

func TestWithdrawWithNotes_InvalidAmount(t *testing.T) {
	acct := account.NewAccount(1000)

	_, err := atm.WithdrawWithNotes(25, acct)
	if err == nil {
		t.Fatal("expected error, received nil")
	}
}

func TestWithdrawWithNotes_InsufficientBalance(t *testing.T) {
	acct := account.NewAccount(100)

	_, err := atm.WithdrawWithNotes(200, acct)
	if err == nil {
		t.Fatal("expected error, received nil")
	}
}

func TestWithdrawWithNotes_WithdrawAllBalance(t *testing.T) {
	acct := account.NewAccount(130)

	notes, err := atm.WithdrawWithNotes(130, acct)
	if err != nil {
		t.Fatalf("error to withdraw: %v", err)
	}

	expected := []atm.Notes{
		{Note: 100, Quantity: 1},
		{Note: 20, Quantity: 1},
		{Note: 10, Quantity: 1},
	}

	if !reflect.DeepEqual(notes, expected) {
		t.Errorf("expected %v, received %v", expected, notes)
	}

	if acct.GetBalance() != 0 {
		t.Errorf("balance should be 0: %d", acct.GetBalance())
	}
}
