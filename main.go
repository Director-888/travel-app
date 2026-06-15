package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		result := 5 * i
		fmt.Println("5*", i, "=", result)
	}
}
