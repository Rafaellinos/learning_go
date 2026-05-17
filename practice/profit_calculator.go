package main

import "fmt"

func calculateEbt(revenue, expenses float64) (result float64) {
	return revenue - expenses
}

func calculateProfit(ebt, taxRate float64) float64 {
	return ebt * (1 - taxRate/100)
}

func calculateRatio(ebt, profit float64) float64 {
	return ebt / profit
}

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("tax rate: ")
	fmt.Scan(&taxRate)

	ebt := calculateEbt(revenue, expenses)
	profit := calculateProfit(ebt, taxRate)
	ratio := calculateRatio(ebt, profit)

	fmt.Printf("EBT: %.2f\n", ebt)
	fmt.Println(fmt.Sprintf("After tax: %.2f", profit))
	fmt.Println(fmt.Sprintf("Ratio: %.2f", ratio))
    // test
}

