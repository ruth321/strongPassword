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

	if len(pw) < 6 || checkSpecChar(pw) ||  {

	}

}

func checkSpecChar(str string) (ch bool) {
	for i := 0; i < len(str); i++ {
		if !strings.Contains("!@#$%^&*()-+_", string(str[i])) {
			ch = false
			return ch
		}
	}
	ch = true
	return ch
}
