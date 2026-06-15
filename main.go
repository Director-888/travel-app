package main

import "fmt"

func main() {

	var correctpassword string
	fmt.Println("Введите пароль")
	fmt.Scanln(&correctpassword)

	password := "secret123"

	if correctpassword == password {
		fmt.Println("Вход успешно выполнен")

	} else {
		fmt.Println("Неверный пароль, доступ заблокирован")
	}
}
