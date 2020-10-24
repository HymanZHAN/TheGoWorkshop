package main

import "fmt"

func main() {
	printHiName("305")
}

func printHiName(key string) {
	nameMap := map[string]string{
		"305": "Sue",
		"204": "Bob",
		"631": "Jake",
		"073": "Tracy",
	}
	fmt.Println("Hi,", nameMap[key])
}