package main

import "fmt"

//Дан слайс []int{-3, 5, -1, 8, 0, -7, 4}.
//Создай новый слайс только из положительных чисел (> 0).

func main() {
	nums := []int{-3, 5, -1, 8, 0, -7, 4}
	new_slice := []int{}

	for _, v := range nums {
		if v > 0 {
			new_slice = append(new_slice, v)
		}
	}
	fmt.Println("Положительные:", new_slice)
}
