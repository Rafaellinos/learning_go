package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/Rafaellinos/bank/fileops"
)

const balanceFileName = "balance.txt"

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


func main() {

	accountBalance, err := fileops.GetDecimalFromFile(balanceFileName)
	if err != nil {
		panic(err)
	}

	fmt.Println("Wecome to Go Bank!")
	showMenu()
	var choice int // default is 0
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
				fileops.WriteDecimalToFile(balanceFileName, accountBalance)
				fmt.Printf("Your new Balance is: %v\n", accountBalance)
			case 3:
				fmt.Print("Value to withdraw: ")
				var valueWithdraw float64
				fmt.Scan(&valueWithdraw)
				if !checkAmount(accountBalance, valueWithdraw) {
					fmt.Println("Current balance is: ", accountBalance)
				} else {
					accountBalance = accountBalance.Sub(decimal.NewFromFloat(valueWithdraw))
					fileops.WriteDecimalToFile(balanceFileName, accountBalance)
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


