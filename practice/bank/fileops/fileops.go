package fileops

import (
	"os"
	"errors"
	"github.com/shopspring/decimal"
)


func WriteDecimalToFile(fileName string, value decimal.Decimal) {
	balanceText := value.String()
	os.WriteFile(fileName, []byte(balanceText), 0644)
}

func GetDecimalFromFile(fileName string) (decimal.Decimal, error) {
	// _ is ignored by go compiler
	if _, err := os.Stat(fileName); err != nil {
		defaultAmount, _ := decimal.NewFromString("0")
		WriteDecimalToFile(fileName, defaultAmount)
		return defaultAmount, nil
	}

	valueByte, err := os.ReadFile(fileName)
	if err != nil {
		return decimal.Decimal{}, errors.New("failed to read file")
	}
	valueDecimal, err := decimal.NewFromString(string(valueByte))
	if err != nil {
		return decimal.Decimal{}, errors.New("failed to parse value from file")
	}
	return valueDecimal, nil
}

