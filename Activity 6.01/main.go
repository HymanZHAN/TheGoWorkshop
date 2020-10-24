package main

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidLastName  = errors.New("invalid last name")
	ErrInvalidFirstName = errors.New("invalid first name")
)

func main() {
	fmt.Println(ErrInvalidFirstName.Error())
	fmt.Println(ErrInvalidLastName)
}
