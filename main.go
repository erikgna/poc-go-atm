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

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter amount to withdraw (0 to exit): ")

		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Invalid amount")
			continue
		}

		amount, err := strconv.Atoi(text[:len(text)-1])
		if err != nil {
			fmt.Println("Invalid amount")
			continue
		}

		if amount == 0 {
			fmt.Println("Exiting...")
			break
		}

		account := account.NewAccount(999999999)
		notes, err := atm.WithdrawWithNotes(amount, account)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(notes)
	}
}
