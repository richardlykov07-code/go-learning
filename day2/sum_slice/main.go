package main

import "fmt"

func main() {
	nums := []int{3, 7, 2, 9, 5}
	total := 0
	for _, v := range nums {
		total += v
	}
	fmt.Println("Сумма:", total)
}
