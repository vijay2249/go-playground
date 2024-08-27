package files

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const balanceFileName = "balance.txt"

func WriteBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(balanceFileName, []byte(balanceText), 0644)
}

func GetBalanceToFile() (float64, error) {
	data, err := os.ReadFile(balanceFileName)
	if err != nil {
		return 0, errors.New("failed to find and read file")
	}
	balance, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return 0, errors.New("failed to convert data to string")
	}
	return balance, nil
}

func WriteDataToFile(fileName string, data string) {
	os.WriteFile(fileName, []byte(data), 0644)
}
