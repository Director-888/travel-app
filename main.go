package main

import "fmt"

func main() {
	var totalPrice int

	ticketPrice := 10000

	var age int
	fmt.Println("Сколько вам лет?")
	fmt.Scanln(&age)

	if age < 7 {
		fmt.Println("У вас скидка 50%!")
		totalPrice = ticketPrice - (ticketPrice / 2)

	} else if age > 65 {
		fmt.Println("У вас скидка 30%")
		totalPrice = ticketPrice - (ticketPrice * 30 / 100)
	} else {
		totalPrice = ticketPrice
	}
	fmt.Println("Стоимость вашего билета", totalPrice, "рублей")

}
