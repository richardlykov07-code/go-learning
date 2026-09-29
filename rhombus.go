package main

import "fmt"

func rhombus() {
	var n int
	fmt.Print("Введите нечетное число(высоту ромба):")
	fmt.Scan(&n)
	//верхняя граница
	for i := 1; i <= (n+1)/2; i++ {
		for j := 0; j < (n+1)/2-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	//нижняя граница
	for i := (n+1)/2 - 1; i >= 1; i-- {
		for j := 0; j < (n+1)/2-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	rhombus()
}
