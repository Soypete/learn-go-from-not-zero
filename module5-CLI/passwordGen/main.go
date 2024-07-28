package main

import (
	"flag"
	"math/rand"
	"os"
)

var requiredChars string
var length int

func main() {
	flag.StringVar(&requiredChars, "required", "&$*", "Required characters")
	flag.IntVar(&length, "length", 8, "Length of the password")
	flag.Parse()

	password := generatePassword([]byte(requiredChars), length)
	os.Stdout.WriteString(password + "\n")
}

var alpha = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func generatePassword(requiredChars []byte, length int) string {
	password := make([]byte, length)
	// Your code here
	for range length {
		i := rand.Intn(len(alpha) - 1)
		password = append(password, alpha[i])
	}
	for _, v := range requiredChars {
		r := rand.Intn(len(password))
		password[r] = v
	}
	return string(password)
}
