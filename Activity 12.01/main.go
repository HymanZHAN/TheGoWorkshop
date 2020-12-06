package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var ErrBudgetCategoryNotFound = errors.New("budget category is not found")

type budgetCategory string

type transaction struct {
	ID       int
	payee    string
	spent    float64
	category budgetCategory
}

var categoryMap = map[budgetCategory]budgetCategory{
	"fuel":           "autoFuel",
	"gas":            "autoFuel",
	"food":           "food",
	"mortgage":       "mortgage",
	"repairs":        "repairs",
	"life insurance": "insurance",
	"car insurance":  "insurance",
	"utilities":      "utilities",
}

func main() {
	transactionFile := flag.String("f", "", "Location of the transaction csv file")
	logFile := flag.String("l", "errors.log", "Location of error logs")

	flag.Parse()

	if *transactionFile == "" {
		log.Fatal("Must provide path to the bank transaction file.")
	}

	transF, err := os.Open(*transactionFile)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = transF.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	_, err = os.Stat(*logFile)
	if !os.IsNotExist(err) {
		os.Remove(*logFile)
	}

	transactions := parseBankFile(transF, *logFile)
	// parseBankFile(transF, *logFile)

	for _, transaction := range transactions {
		fmt.Println(transaction)
	}

}

func convertToBudgetCategory(category budgetCategory) (budgetCategory, error) {
	mappedCategory, ok := categoryMap[category]
	if !ok {
		return "", ErrBudgetCategoryNotFound
	}
	return mappedCategory, nil
}

func parseBankFile(bankTransactions io.Reader, logFile string) []transaction {
	r := csv.NewReader(bankTransactions)
	header := true
	var transactions []transaction
	for {
		var trans transaction
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if !header {
			for idx, value := range record {
				value = strings.Trim(value, " ")
				switch idx {
				case 0: // id
					id, err := strconv.Atoi(value)
					if err != nil {
						writeErrorToLog("ID not a valid int", err, value, logFile)
					}
					trans.ID = id
				case 1: // payee
					trans.payee = value
				case 2: // spent
					spent, err := strconv.ParseFloat(value, 64)
					if err != nil {
						writeErrorToLog("spent not a valid float64", err, value, logFile)
					}
					trans.spent = spent
				case 3: // category
					mappedCategory, err := convertToBudgetCategory(budgetCategory(value))
					if err != nil {
						writeErrorToLog("Category not found", err, value, logFile)
					}
					trans.category = mappedCategory
				}

			}
			transactions = append(transactions, trans)
		}
		header = false
	}

	return transactions
}

func writeErrorToLog(msg string, err error, data string, logfile string) {
	logF, err := os.OpenFile(logfile, os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = logF.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	errMsg := fmt.Sprintf("Msg: %s | Err: %v | Data: %s\n", msg, err, data)
	logF.WriteString(errMsg)
}
