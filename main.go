package main

import "fmt"

func main() {
	var name string
	fmt.Println("Как тебя зовут?")
	fmt.Scanln(&name)

	var age int
	fmt.Println("Сколько тебе лет?")
	fmt.Scanln(&age)

	var city string
	fmt.Println("В каком городе живешь?")
	fmt.Scanln(&city)

	fmt.Println("Привет", name, "Тебе", age, "лет", "и ты живешь в городе", city)
}
