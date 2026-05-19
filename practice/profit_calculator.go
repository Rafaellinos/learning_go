package main

import (
	"fmt"
	"os"
	"errors"
	"time"
	"strings"
	"strconv"
	"bufio"
)

const calcFileName = "revenue_calc.csv"
const csvHeaders = "EBT,Profit,Ratio,Time"

func convertToFloat(value string) float64 {
	valueFloat, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic(err)
	}
	return valueFloat
}

func getCalcFromFile() (float64, float64, float64) {
	if _, err := os.Stat(calcFileName); err != nil {
		writeHeader()
		return 0,0,0
	}
	file, err := os.Open(calcFileName)
	if err != nil {
		fmt.Print(err)
		return 0,0,0
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lastLine string
	
	for scanner.Scan() {
		lastLine = scanner.Text()
	}
	if strings.Contains(lastLine, "EBT") {
		return 0,0,0
	}
	valuesArr := strings.Split(lastLine, ",")
	return convertToFloat(valuesArr[0]), convertToFloat(valuesArr[1]), convertToFloat(valuesArr[2])
}

func writeHeader() {
	os.WriteFile(calcFileName, []byte(csvHeaders+"\n"), 0644)
}

func appendCalc(ebt, profit, ratio float64) {
	// append and write to a file
	file, err := os.OpenFile(calcFileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()

	line := fmt.Sprintf("%.2f,%.2f,%.2f,%v\n", ebt, profit, ratio, time.Now().Format(time.DateTime))
	if _, err := file.WriteString(line); err != nil {
		fmt.Println(err)
	}
}


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

	ebtR, profitR, ratioR := getCalcFromFile()
	fmt.Println("Last calc: ")
	fmt.Printf("EBT: %.2f", ebtR)
	fmt.Printf("Profit: %.2f", profitR)
	fmt.Printf("Ratio: %.2f\n", ratioR)
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("revenue: ")
	fmt.Scan(&revenue)

	if revenue <= 0 {
		panic(errors.New("invalid revenue value, should be greater than 0"))
	}

	fmt.Print("expenses: ")
	fmt.Scan(&expenses)

	if expenses < 0 {
		panic(errors.New("invalid expenses value, shoud be greater or equal to 0"))
	}

	fmt.Print("tax rate: ")
	fmt.Scan(&taxRate)

	if taxRate < 0 {
		panic(errors.New("invalid taxRatem should be greater or equal to 0"))
	}

	ebt := calculateEbt(revenue, expenses)
	profit := calculateProfit(ebt, taxRate)
	ratio := calculateRatio(ebt, profit)

	fmt.Printf("EBT: %.2f\n", ebt)
	fmt.Println(fmt.Sprintf("After tax: %.2f", profit))
	fmt.Println(fmt.Sprintf("Ratio: %.2f", ratio))

	appendCalc(ebt, profit, ratio)

}

