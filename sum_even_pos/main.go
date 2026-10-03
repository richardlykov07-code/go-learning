package main

import "fmt"

func main() {
	nums := []int{10, 20, 30, 40, 50, 60}
	sum := 0
	for i, v := range nums {
		if i%2 == 0 {
			sum += v
		}
	}
	fmt.Println("Сумма на чётных позициях:", sum)
}
