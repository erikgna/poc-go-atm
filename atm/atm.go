package atm

import (
	"errors"
	"go-atm/account"
)

type Notes struct {
	Note     int
	Quantity int
}

var denominations = []int{100, 50, 20, 10}

func WithdrawWithNotes(amount int, account *account.Account) ([]Notes, error) {
	if amount%10 != 0 {
		return nil, errors.New("amount is not divisible by 10")
	}

	if !account.HasEnoughBalance(amount) {
		return nil, errors.New("insufficient balance")
	}

	err := account.Withdraw(amount)
	if err != nil {
		return nil, err
	}

	notes := []Notes{}
	for _, denom := range denominations {
		count := amount / denom
		if count > 0 {
			notes = append(notes, Notes{Note: denom, Quantity: count})
			amount -= denom * count
		}
	}

	return notes, nil
}
