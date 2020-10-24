package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type locale struct {
	language string
	country  string
}

func main() {
	var locales []locale
	locales = append(locales, locale{"en", "US"})
	locales = append(locales, locale{"en", "CN"})
	locales = append(locales, locale{"fr", "CN"})
	locales = append(locales, locale{"fr", "FR"})
	locales = append(locales, locale{"ru", "RU"})

	args, err := getPassedArgs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	arg := args[0]

	var passedLocale locale
	var isSupported bool
	passedLocale.language = strings.Split(arg, "_")[0]
	passedLocale.country = strings.Split(arg, "_")[1]

	for _, l := range locales {
		if l.country == passedLocale.country && l.language == passedLocale.language {
			fmt.Println("Locale passed is supported")
			isSupported = true
		}
	}

	if !isSupported {
		fmt.Println("Locale not supported:", arg)
	}

}

func getPassedArgs() ([]string, error) {
	var args []string
	if len(os.Args) <= 1 {
		return args, errors.New("No arg passed")
	}
	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}
	return args, nil
}

func getLocales(extraLocals []string) []string {
	var locales []string
	locales = append(locales, "en_US", "fr_FR")
	locales = append(locales, extraLocals...)
	return locales
}

func findLongest(args []string) string {
	var longest string

	for i := 0; i < len(args); i++ {
		if len(longest) < len(os.Args[i]) {
			longest = args[i]
		}
	}

	return longest
}
