package main

import "fmt"

func main() {
	var days int
	fmt.Println("Сколько дней продлится поездка?")
	fmt.Scanln(&days)

	hotelPrice := 4000
	for i := 1; i <= days; i++ {
		currentTotal := i * hotelPrice
		fmt.Println("День", i, "стоимость", currentTotal, "рублей")
	}
}
