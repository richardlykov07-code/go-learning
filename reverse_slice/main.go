package main

import "fmt"

func reverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		j := len(nums) - i - 1
		nums[i], nums[j] = nums[j], nums[i]
	}
	fmt.Println("Перевёрнутый слайс:", nums)
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Исходный слайс:", slice)
	reverse(slice)
}
