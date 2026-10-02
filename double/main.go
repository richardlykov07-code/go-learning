package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	result := make([]int, 0, len(nums))
	for _, v := range nums {
		result = append(result, v*2)
	}
	fmt.Println("Исходный:", nums)
	fmt.Println("Удвоенный:", result)
}
