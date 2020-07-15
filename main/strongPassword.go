package main

import (
	"fmt"
	"github.com/labstack/gommon/random"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var pw string
	fmt.Print("Enter password: ")
	fmt.Scan(&pw)

	pwEx := []byte(pw)

	for true {
		if len(pw) > 5 && checkSpecChar(pw) && checkUpCase(pw) && checkLowCase(pw) && checkDigit(pw) {
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
				fmt.Printf(" - at least 1 special character \"%s\"\n", random.Symbols)
				pwEx = append(pwEx, random.Symbols[rand.Intn(len(random.Symbols))])
			}
			if !checkUpCase(pw) {
				fmt.Println(" - at least 1 uppercase English character")
				pwEx = append(pwEx, byte(65+rand.Intn(26)))
			}
			if !checkLowCase(pw) {
				fmt.Println(" - at least 1 lowercase English character")
				pwEx = append(pwEx, byte(97+rand.Intn(26)))
			}
			if !checkDigit(pw) {
				fmt.Println(" - at least 1 digit")
				pwEx = append(pwEx, byte(48+rand.Intn(10)))
			}
			if len(pwEx) < 6 {
				for i, g := 0, len(pwEx); i < 6-g; i++ {
					pwEx = append(pwEx, byte(97+rand.Intn(26)))
				}
			}

			fmt.Println("Example:", string(pwEx))

			fmt.Print("\nEnter password: ")
			fmt.Scan(&pw)
			fmt.Println()

			pwEx = []byte(pw)
		}
	}
}

func checkSpecChar(s string) (ch bool) {
	for i := 0; i < len(s); i++ {
		if strings.Contains(random.Symbols, string(s[i])) {
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
