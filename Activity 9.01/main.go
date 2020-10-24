package main

import (
	"errors"
	"log"
	"strings"
	"unicode"
)

var ErrInvalidSSNLength error = errors.New("ssn is not nine characters long")
var ErrInvalidSSNNumbers error = errors.New("ssn has non-numeric digits")
var ErrInvalidSSNPrefix error = errors.New("ssn has three zeros as a prefix")
var ErrInvalidSSNDigitPlace error = errors.New("ssn starts with a 9 requires 7 or 9 in the fourth place")

func main() {
	validateSSN := []string{
		"123-45-6789",
		"012-8-678",
		"000-12-0962",
		"999-33-3333",
		"087-65-4321",
		"123-45-zzzz",
	}
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Printf("Checking data: %#v\n", validateSSN)

	for i, ssn := range validateSSN {
		log.Printf("Validate data \"%s\" %d of %d", ssn, i+1, len(validateSSN))
		newSsn := strings.Replace(ssn, "-", "", -1)
		var err error

		err = checkSSNLength(newSsn)
		if err != nil {
			log.Printf("the value of %s caused an error: %s\n", ssn, err)
		}
		err = checkSSNDigitPlace(newSsn)
		if err != nil {
			log.Printf("the value of %s caused an error: %s\n", ssn, err)
		}
		err = checkSSNPrefix(newSsn)
		if err != nil {
			log.Printf("the value of %s caused an error: %s\n", ssn, err)
		}
		err = checkSSNNumbers(newSsn)
		if err != nil {
			log.Printf("the value of %s caused an error: %s\n", ssn, err)
		}
	}
}

func checkSSNLength(ssn string) error {
	if len(ssn) != 9 {
		return ErrInvalidSSNLength
	}
	return nil
}

func checkSSNNumbers(ssn string) error {
	for _, digit := range ssn {
		if !unicode.IsNumber(digit) {
			return ErrInvalidSSNNumbers
		}
	}
	return nil
}

func checkSSNPrefix(ssn string) error {
	if strings.HasPrefix(ssn, "000") {
		return ErrInvalidSSNPrefix
	}
	return nil
}

func checkSSNDigitPlace(ssn string) error {
	if strings.HasPrefix(ssn, "9") {
		if ssn[3] != '7' && ssn[3] != '9' {
			return ErrInvalidSSNDigitPlace
		}
	}
	return nil
}
