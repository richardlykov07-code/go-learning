package main

import "fmt"

func main() {
	chet_count := 0
	nechet_count := 0
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, v := range nums {
		if v%2 == 0 {
			chet_count++
		} else {
			nechet_count++
		}
	}
	fmt.Println("Чётные:", chet_count)
	fmt.Println("Нечётные:", nechet_count)
}
