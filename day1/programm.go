package main

import (
	"fmt"
	"math/rand/v2"
)

// функция бесконечного цикла
func commandLoop() {
	var input string
	for input != "exit" {
		fmt.Print("Введите команду: ")
		fmt.Scan(&input)
	}
	fmt.Println("Программа завершена.")
}

// альтернативный вариант бесконечного цикла
func whileTrue() {
	for {
		var input string
		fmt.Print("Введите команду: ")
		fmt.Scanln(&input)
		if input == "exit" {
			break
		}
	}
	fmt.Println("Программа завершена.")
}

// сумма чисел
func SumNum(num int) int {
	sum := 0

	for i := 1; i <= num; i++ {
		sum += i
	}
	return sum
}

// таблица умножения
func multable() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			fmt.Printf("%4d", i*j)
		}
		fmt.Println()
	}
}

// Угадай число
func guess() {
	fmt.Println("Я загадал число от 1 до 100")
	secret := rand.IntN(100) + 1
	var n int
	count := 0
	for {
		fmt.Print("Ваш вариант:")
		fmt.Scan(&n)
		count++
		if n == secret {
			fmt.Printf("Угадали! Число было %d. Попыток: %d \n ", secret, count)
			break
		} else if n >= secret {
			fmt.Println("Меньше!")
		} else if n <= secret {
			fmt.Println("Больше!")
		}

	}
}

// ромб
func rhombus() {
	var N int
	fmt.Print("Введите нечётное число:")
	fmt.Scanln(&N)

	for i := 1; i < N; i++ {
		for j := 1; j < (N+1)/2; j++ {
			fmt.Print(" ")
		}
		for j := 1; j < (N+1)/2; j++ {
			fmt.Print("*")
		}
	}
}

func FizzBuzz() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%5 == 0:
			fmt.Println("Buzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		default:
			fmt.Println(i)
		}
	}
}

func main() {
	/*
		var n int
		fmt.Print("Введите N: ")
		fmt.Scanln(&n)
		result := SumNum(n)
		fmt.Println("Сумма:", result)

		fmt.Println("Таблица умножения:")
		multable()

		guess()
		FizzBuzz()
	*/

	rhombus()
}
