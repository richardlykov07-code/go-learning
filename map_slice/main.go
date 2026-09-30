package main

import "fmt"

// Программа спрашивает N. Считает и выводит сумму только чётных чисел от 1 до N.
func even_sum(n int) int {
	sum := 0
	for i := 2; i <= n; i += 2 {
		sum += i
	}
	return sum
}

func mult_n() {
	var n int
	fmt.Print("Введите N:")
	fmt.Scan(&n)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d * %d = %d\n", n, i, i*n)
	}
}

func digit_sum() {
	var n int
	fmt.Print("Введите число:")
	fmt.Scan(&n)
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit
		n = n / 10
	}
	fmt.Print("Сумма:", sum)
}

//Программа спрашивает целое число. Выводит его в обратном порядке цифр.

func reverse_num() {
	var n int
	fmt.Print("Введите число:")
	fmt.Scan(&n)
	result := 0
	for n > 0 {
		result = result*10 + n%10
		n /= 10
	}
	fmt.Print("Развёрнутое: ", result)
}

//Программа спрашивает слово.
//Считает и выводит количество гласных букв (а, е, ё, и, о, у, ы, э, ю, я).

func vowels() {
	var word string
	count := 0
	fmt.Print("Введите слово:")
	fmt.Scan(&word)
	for _, r := range word {
		switch r {
		case 'а', 'е', 'ё', 'и', 'о', 'у', 'ы', 'э', 'ю', 'я':
			count++
		}
	}
	fmt.Print("Гласных:", count)
}

//Создай слайс целых чисел []int{5, 10, 15, 20, 25} прямо в коде.
//Посчитай сумму и среднее через for range.

func slice_sum() {
	total := 0
	sum := []int{5, 10, 15, 20, 25}
	for _, v := range sum {
		total += v
	}
	srd := float64(total) / float64(len(sum))
	fmt.Printf("Сумма:  %.2f ", float64(total))
	fmt.Println()
	fmt.Printf("Среднее: %.2f ", float64(srd))
}

func main() {
	var n int
	fmt.Print("Введите N:")
	fmt.Scan(&n)
	fmt.Print("Сумма чётных:", even_sum(n))

	mult_n()

	digit_sum()

	reverse_num()

	vowels()

	slice_sum()
}
