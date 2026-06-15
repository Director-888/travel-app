package main

import "fmt"

func main() {

	var age int
	fmt.Println("Сколько тебе лет?")
	fmt.Scanln(&age)

	if age >= 18 {
		fmt.Println("Доступ разрешен!")

	} else {
		fmt.Println("Доступ запрещен, вам нужно подрасти!")
	}
}
