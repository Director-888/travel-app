package main

import "fmt"

func main() {
	var days int
	fmt.Println("Сколько дней продлится поездка?")
	fmt.Scanln(&days)

	hotelPrice := 4000
	total := hotelPrice * days

	// Проверяем условие: если дней больше 7
	if days > 14 {
		fmt.Println("Ура! Вы получили скидку 20% ")
		total = total - (total / 5)
	} else if days > 7 {
		fmt.Println("Ура! Вы получили скидку 10% ")
		total = total - (total / 10)
	} else if days < 7 {
		fmt.Println("К сожалению у вас нет скидки")

	}

	fmt.Println("Итоговая стоимость отеля обойдется в", total, "рублей")
}
