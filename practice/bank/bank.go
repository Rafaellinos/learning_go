package main

import (
	"fmt"
	"os"
	"github.com/shopspring/decimal"
)

const balanceFileName = "balance.txt"

func writeBalanceToFile(balance decimal.Decimal) {
	balanceText := balance.String()
	os.WriteFile(balanceFileName, []byte(balanceText), 0644)
}

func getBalanceFromFile() decimal.Decimal {
	// _ is ignored by go compiler
	if _, err := os.Stat(balanceFileName); err != nil {
		defaultAmount, _ := decimal.NewFromString("0")
		writeBalanceToFile(defaultAmount)
		return defaultAmount
	}

	valueByte, err := os.ReadFile(balanceFileName)
	if err != nil {
		panic(err)
	}
	valueDecimal, err := decimal.NewFromString(string(valueByte))
	if err != nil {
		panic(err)
	}
	return valueDecimal
}

func checkAmount(currentAmount decimal.Decimal, withDrawAmount float64) bool {
	if withDrawAmount <= 0 {
		fmt.Println("Withdrawal amount cannot be lower or equal than 0.")
		return false
	}
	if currentAmount.GreaterThanOrEqual(decimal.NewFromFloat(withDrawAmount)) {
		return true
	}
	fmt.Println("Withdrawal amount must be equal or lower than your current balance")
	return false
}

func showMenu() {
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
}

func main() {

	accountBalance := getBalanceFromFile()

	fmt.Println("Wecome to Go Bank!")
	showMenu()
	var choice int
	for choice != 4 {
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		switch choice {
			case 1:
				fmt.Printf("Your Balance is: %v\n", accountBalance)
			case 2:
				fmt.Print("Value to deposit: ")
				var depositValue float64
				fmt.Scan(&depositValue)
				if depositValue <= 0 {
					fmt.Println("invalid ammount. Must be greater than 0")
					continue
				}
				accountBalance = accountBalance.Add(decimal.NewFromFloat(depositValue)) // does not mutate the object, no pointer
				writeBalanceToFile(accountBalance)
				fmt.Printf("Your new Balance is: %v\n", accountBalance)
			case 3:
				fmt.Print("Value to withdraw: ")
				var valueWithdraw float64
				fmt.Scan(&valueWithdraw)
				if !checkAmount(accountBalance, valueWithdraw) {
					fmt.Println("Current balance is: ", accountBalance)
				} else {
					accountBalance = accountBalance.Sub(decimal.NewFromFloat(valueWithdraw))
					writeBalanceToFile(accountBalance)
					fmt.Printf("Your new Balance is: %v\n", accountBalance)
				}
			case 4:
				fmt.Println("Goodbye!")
				return
			default:
				fmt.Println("invalid choice")
		}
		showMenu()
	}
}


