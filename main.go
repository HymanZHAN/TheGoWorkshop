package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	LastName  string  `json:"lname"`
	FirstName string  `json:"fname"`
	Address   address `json:"address"`
}
type address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode int    `json:"zipcode"`
}

func main() {
	p := person{LastName: "Vader", FirstName: "Darth"}
	p.Address = address{
		Street:  "Galaxy Far Away",
		City:    "Dark Side",
		State:   "Tatooine",
		ZipCode: 12345,
	}

	noPrettyPrint, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	prettyPrint, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(noPrettyPrint))
	fmt.Println()
	fmt.Println(string(prettyPrint))
}
