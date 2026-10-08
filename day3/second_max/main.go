package main

import "fmt"

//Дан слайс []int{10, 5, 8, 20, 15}. Найди второй по величине элемент.

func main() {
	nums := []int{10, 5, 8, 20, 15}
	var max, second int
	for _, v := range nums {
		if v > max {
			second = max
			max = v
		} else if v > second && second != max {
			second = v
		}
	}
	fmt.Println("Второй максимум:", second)
}
