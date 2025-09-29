package main

import (
	"bufio"
	"fmt"
	"go-atm/account"
	"go-atm/atm"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter amount: ")

	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid amount")
		return
	}

	amount, err := strconv.Atoi(text[:len(text)-1])
	if err != nil {
		fmt.Println("Invalid amount")
		return
	}

	account := account.NewAccount(999999999)
	notes, err := atm.WithdrawWithNotes(amount, account)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(notes)
}
