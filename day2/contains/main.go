package main

import "fmt"

func main() {
	var number int
	target := []int{4, 8, 15, 16, 23, 42}
	found := false

	fmt.Print("Введите число:")
	fmt.Scan(&number)
	for _, v := range target {
		if v == number {
			found = true
			break
		}
	}

	fmt.Print("Найдено:", found)
}
