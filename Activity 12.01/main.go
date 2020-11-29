package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"io"
	"log"
	"os"
)

var ErrWorkingfileNotFound = errors.New("The budget category is not found")

type transaction struct {
	ID       int
	payee    string
	spent    float64
	category string
}

type budgetCategory map[string][]transaction

func main() {
	transactionFile := flag.String("f", "", "Location of the transaction csv file.")
	logFile := flag.String("o", "", "Location of the output log file.")

	flag.Parse()

	f, err := os.Open(*transactionFile)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

	}

}
