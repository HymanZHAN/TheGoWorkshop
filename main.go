package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	in := `
firstName, lastName, age
Celina, Jones, 18
Cailyn, Henderson, 13
Cayden, Smith, 42
`

	header := true
	r := csv.NewReader(strings.NewReader(in))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if !header {
			for idx, value := range record {
				switch idx {
				case 0:
					fmt.Println("First Name: ", value)
				case 1:
					fmt.Println("Last Name: ", value)
				case 2:
					fmt.Println("Age: ", value)
				}
			}
		}
		header = false
	}
}
