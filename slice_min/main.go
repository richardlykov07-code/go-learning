package main

import "fmt"

func main() {
	nums := []int{42, 17, 89, 3, 56, 91, 12, 7}
	if len(nums) == 0 {
		fmt.Println("Слайс пуст")
		return
	}
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}
	fmt.Print("Минимум: ", min)
}
