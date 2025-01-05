package main

import (
	"flag"
	"fmt"
	"math/rand"
)

var (
	passLength   int
	useSymbols   bool
	useNumbers   bool
	useUppercase bool

	symbols   = "!@#$%^&*()_+{}|:<>?~"
	numbers   = "0123456789"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercase = "abcdefghijklmnopqrstuvwxyz"
)

func main() {

	flag.IntVar(&passLength, "len", 8, "Password length")
	flag.BoolVar(&useSymbols, "symbols", false, "Use symbols")
	flag.BoolVar(&useNumbers, "numbers", false, "Use numbers")
	flag.BoolVar(&useUppercase, "uppercase", false, "Use uppercase")
	flag.Parse()

	fmt.Println("Generating password...")
	fmt.Println("Password generated: ", generatePassword(passLength, useSymbols, useNumbers, useUppercase))
}

func generatePassword(length int, useSymbols bool, useNumbers bool, useUppercase bool) string {
	password := make([]byte, length)

	// Generate the password
	for i := range length {
		password[i] = lowercase[rand.Intn(len(lowercase)-1)]
	}

	if useSymbols {
		symbolIndex := rand.Intn(len(symbols) - 1)
		passwordIndex := rand.Intn(length - 1)
		password[passwordIndex] = symbols[symbolIndex]
	}

	if useNumbers {
		numberIndex := rand.Intn(len(numbers) - 1)
		passwordIndex := rand.Intn(length - 1)
		password[passwordIndex] = numbers[numberIndex]
	}

	if useUppercase {
		uppercaseIndex := rand.Intn(len(uppercase) - 1)
		passwordIndex := rand.Intn(length - 1)
		password[passwordIndex] = uppercase[uppercaseIndex]
	}
	return string(password)

}
