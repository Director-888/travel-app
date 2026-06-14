package main

import "fmt"

func main() {
	var days int
	fmt.Println("Сколько продлится поездка?")
	fmt.Scanln(&days)
	hotelprice := 4000
	total := hotelprice * days
	fmt.Println(total)
}
