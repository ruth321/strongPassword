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

	//Инициализация и ввод пароля
	var pw string
	fmt.Print("Enter password: ")
	fmt.Scan(&pw)

	//Инициализация примера безопасного пароля и присваивание ему значения pw в ASCII коде
	pwEx := []byte(pw)

	//Инициализация цикла для проверки безопасности пароля
	for true {
		//Если пароль является безопасным (содержит больше 5 символов, хотя бы 1 специальный символ,
		//строчную и заглавную буквы английского алфавита, и цифру),
		//то выводится соответствующее сообщение и происходит выход из цикла
		if len(pw) > 5 && checkSpecChar(pw) && checkUpCase(pw) && checkLowCase(pw) && checkDigit(pw) {
			fmt.Println("Your password is strong")
			break
		} else {
			//Если пароль не является безопасным, выводится соответствующее сообщение
			fmt.Println("Your password isn't strong.\nYou should add: ")
			//Если пароль меньше необходимой длины, на экран выводится
			//количесво недостающих символов
			if len(pw) < 6 {
				fmt.Print(" - at least ", 6-len(pw))
				if 6-len(pw) > 1 {
					fmt.Println(" symbols")
				} else {
					fmt.Println(" symbol")
				}
			}
			//Далее  информация обо всех недостающих символах, если они имеются, выводится на экран
			//и в pwEx добавляется соответствующий случайный символ
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

			//Если pwEx меньше необходимой длины, к нему добавляются недостающие случайные символы
			//из списка строчных букв английского алфавита
			if len(pwEx) < 6 {
				for i, g := 0, len(pwEx); i < 6-g; i++ {
					pwEx = append(pwEx, byte(97+rand.Intn(26)))
				}
			}

			//Вывод на экран примера безопасного пароля
			fmt.Println("Example:", string(pwEx))

			//Ввод нового пароля
			fmt.Print("\nEnter password: ")
			fmt.Scan(&pw)
			fmt.Println()

			//Присваивание примеру пароля значения pw в ASCII коде
			pwEx = []byte(pw)
		}
	}
}

//Функция проверяет строку s на наличие хотя бы 1 специального символа
//Если он есть, функция возвращает значение true
func checkSpecChar(s string) (b bool) {
	for i := 0; i < len(s); i++ {
		if strings.Contains(random.Symbols, string(s[i])) {
			b = true
			return b
		}
	}
	b = false
	return b
}

//Функция проверяет строку s на наличие хотя бы 1 заглавной буквы из английского алфавита
//Если она есть, функция возвращает значение true
func checkUpCase(s string) (b bool) {
	for i := 0; i < len(s); i++ {
		if s[i] > 64 && s[i] < 91 {
			b = true
			return b
		}
	}
	b = false
	return b
}

//Функция проверяет строку s на наличие хотя бы 1 строчной буквы из английского алфавита
//Если она есть, функция возвращает значение true
func checkLowCase(s string) (b bool) {
	for i := 0; i < len(s); i++ {
		if s[i] > 96 && s[i] < 123 {
			b = true
			return b
		}
	}
	b = false
	return b
}

//Функция проверяет строку s на наличие хотя бы 1 цифры.
//Если она есть, функция возвращает значение true
func checkDigit(s string) (b bool) {
	for i := 0; i < len(s); i++ {
		if strings.Contains("0123456789", string(s[i])) {
			b = true
			return b
		}
	}
	b = false
	return b
}
