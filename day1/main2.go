package main

import "fmt"

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("Введите первое число:")
	fmt.Scan(&num1)

	fmt.Print("Введите второе число:")
	fmt.Scan(&num2)

	fmt.Print("Выберите оператор(+,-,*,/):")
	fmt.Scan(&operator)

	switch operator {
	case "+":
		fmt.Printf("Значение: %.2f ", num1+num2)
	case "-":
		fmt.Printf("Значение: %.2f ", num1-num2)
	case "*":
		fmt.Printf("Значение: %.2f ", num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Ошибка: Деление на ноль ")
		} else {
			fmt.Printf("Значение: %.2f ", num1/num2)
		}

	default:
		fmt.Println("Неизвестный оператор", operator)
	}
}
