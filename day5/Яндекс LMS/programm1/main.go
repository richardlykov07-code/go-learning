package main

import "fmt"

func main() {
	var name string
	fmt.Print("Введите своё имя:")
	fmt.Scanln(&name)

	messageResult := fmt.Sprintf("Привет, %s !", name)
	fmt.Print(messageResult)
}
