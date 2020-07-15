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
	fxPw := make([]byte, 6)

	for true {

		if len(pw) > 5 && checkSpecChar(pw) {
			fmt.Println("Your password is strong")
			//break
		} else {
			fmt.Println("Your password isn't strong.\nYou should add at least 1")
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
