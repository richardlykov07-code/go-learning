package main

import "fmt"

//Дан слайс []int{1, 2, 3, 4, 5}.
//Сдвинь все элементы влево на 1 позицию. Первый элемент уходит в конец.

func main() {
	var K int
	nums := []int{1, 2, 3, 4, 5}
	fmt.Print("Введите К: ")
	fmt.Scan(&K)
	K = K % len(nums)
	for shift := 0; shift < K; shift++ {
		first := nums[0]
		for i := 0; i < len(nums)-1; i++ {
			nums[i] = nums[i+1]
		}
		nums[len(nums)-1] = first
	}
	fmt.Println("Сдвинутый:", nums)
}
