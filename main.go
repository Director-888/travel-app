package main

import "fmt"

func main() {
	var days int
	fmt.Println("Сколько дней продлится поездка?")
	fmt.Scanln(&days)

	hotelPrice := 4000
	total := hotelPrice * days

	// Проверяем условие: если дней больше 7
	if days > 7 {
		fmt.Println("Ура! Вы получили скидку 10% за долгое бронирование!")
		// Рассчитываем скидку (делим общую сумму на 10 и вычитаем её из total)
		total = total - (total / 10)
	} else {
		fmt.Println("До скидки вам не хватило еще немного дней!")

	}

	fmt.Println("Итоговая стоимость отеля обойдется в", total, "рублей")
}
