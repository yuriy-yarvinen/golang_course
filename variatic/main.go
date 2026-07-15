package main

import "fmt"

func main() {

	sumOfNumbers := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)
	println("Sum:", sumOfNumbers)

	numbers := []int{1, 2, 3, 4, 5}
	sumOfNumbersInSlice := sum(numbers...)
	println("Sum of numbers:", sumOfNumbersInSlice)

	const (
		_ = iota
		a
		b
		c
		d
	)

	fmt.Printf("%d \t %b \n ", 1, 1)
	fmt.Printf("%d \t %b \n ", 1<<a, 1<<a)
	fmt.Printf("%d \t %b \n ", 1<<b, 1<<b)
	fmt.Printf("%d \t %b \n ", 1<<c, 1<<c)
	fmt.Printf("%d \t %b \n ", 1<<d, 1<<d)
	fmt.Printf("%d \t %b \n ", 3, 3)
	fmt.Printf("%d \t %b \n ", 1<<50, 1<<50)
	fmt.Printf("%d \t %b \n ", 1125899906842625, 1125899906842625)

	type ByteSize int

	const (
		_           = iota // ignore first value by assigning to blank identifier
		KB ByteSize = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
	)

	fmt.Printf("%d \t\t\t %b\n", KB, KB)
	fmt.Printf("%d \t\t %b\n", MB, MB)
	fmt.Printf("%d \t\t %b\n", GB, GB)
	fmt.Printf("%d \t\t %b\n", TB, TB)
	fmt.Printf("%d \t %b\n", PB, PB)
	fmt.Printf("%d \t %b\n", EB, EB)
}

func sum(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
