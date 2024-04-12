package main

import (
	"fmt"
	"log"
	"math/rand"
)

var gueNumber int

func main() {
	fmt.Println("Угадай число и получи приз")
	min, max := 1, 50
	secretNumber := rand.Intn(max-min) + min

	for {
		fmt.Print("Введите ваше предположение: ")
		_, err := fmt.Scanln(&gueNumber)
		if err != nil {
			log.Fatal(err)
		}

		if gueNumber == secretNumber {
			fmt.Println("Поздравляем, вы угадали число!Ваш приз https://www.youtube.com/watch?v=wjI_iJyjiU8&pp=ygUM0YDQuNC60YDQvtC7")
			break
		} else {
			fmt.Println("Увы, цифры были", secretNumber, "поощерительный приз https://www.youtube.com/watch?v=wjI_iJyjiU8&pp=ygUM0YDQuNC60YDQvtC7 ")
		}
	}
}
func fore() {
	for secretNumber := 0; secretNumber < 25; secretNumber++ {
		if gueNumber == secretNumber {
			break
		}
	}
}
