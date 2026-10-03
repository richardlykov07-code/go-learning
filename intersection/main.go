package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{4, 5, 6, 7, 8}
	section := make(map[int]bool)
	result := []int{}
	added := make(map[int]bool)

	for _, v := range b {
		section[v] = true
	}

	for _, v := range a {
		if section[v] == true && added[v] == false {
			result = append(result, v)
			added[v] = true
		}
	}
	fmt.Print(result)
}