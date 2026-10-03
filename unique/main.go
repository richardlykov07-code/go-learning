package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 3, 3, 3, 4, 5, 5}
	result := []int{}
	seen := make(map[int]bool)

	for _, v := range nums {
		if !seen[v] {
			result = append(result, v)
			seen[v] = true
		}
	}
	fmt.Println("Исходный: ", nums)
	fmt.Println("Без дубликатов: ", result)
}
