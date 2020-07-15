package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var pw string
	fmt.Print("Enter password: ")
	fmt.Scan(&pw)
	fmt.Println()

	fxPw := []byte(pw)

	for true {

		if len(pw) > 5 && checkSpecChar(pw) {
			fmt.Println("Your password is strong")
			break
		} else {
			fmt.Println("Your password isn't strong.\nYou should add: ")
			if len(pw) < 6 {
				fmt.Print(" - at least ", 6-len(pw))
				if 6-len(pw) > 1 {
					fmt.Println(" symbols")
				} else {
					fmt.Println(" symbol")
				}

			}
			if !checkSpecChar(pw) {
				fmt.Println(" - at least 1 special character \"!@#$%^&*()-+_\"")

			}
			if !checkUpCase(pw) {
				fmt.Println(" - at least 1 uppercase English character")
			}
			if !checkLowCase(pw) {
				fmt.Println(" - at least 1 lowercase English character")
			}
			if !checkDigit(pw) {
				fmt.Println(" - at least 1 digit")
			}

			break
		}
	}
}

func checkSpecChar(s string) (ch bool) {
	for i := 0; i < len(s); i++ {
		if strings.Contains("!@#$%^&*()-+_", string(s[i])) {
			ch = true
			return ch
		}
	}
	ch = false
	return ch
}

func checkUpCase(s string) (ch bool) {
	for i := 0; i < len(s); i++ {
		if s[i] > 64 && s[i] < 91 {
			ch = true
			return ch
		}
	}
	ch = false
	return ch
}

func checkLowCase(s string) (ch bool) {
	for i := 0; i < len(s); i++ {
		if s[i] > 96 && s[i] < 123 {
			ch = true
			return ch
		}
	}
	ch = false
	return ch
}

func checkDigit(s string) (ch bool) {
	for i := 0; i < len(s); i++ {
		if strings.Contains("0123456789", string(s[i])) {
			ch = true
			return ch
		}
	}
	ch = false
	return ch
}
