package main

import "fmt"

func main() {
	var days int
	fmt.Println("Сколько дней продлится поездка?")
	fmt.Scanln(&days)

	var contry string
	fmt.Println("В какую страну вы едите?")
	fmt.Scanln(&contry)

	hotelPrice := 4000
	total := hotelPrice * days

	// Проверяем условие: если дней больше 7
	if days > 10 && contry == "Грузия" {
		fmt.Println("Ура! Вы получили скидку 20% как VIP клиент!")
		// Рассчитываем скидку (делим общую сумму на 10 и вычитаем её из total)
		total = total - (total / 5)
	} else {
		fmt.Println("Скидка не применяется")

	}

	fmt.Println("Итоговая стоимость отеля обойдется в", total, "рублей")
}
