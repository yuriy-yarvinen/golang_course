package main

func main() {

	sumOfNumbers := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)
	println("Sum:", sumOfNumbers)

	numbers := []int{1, 2, 3, 4, 5}
	sumOfNumbersInSlice := sum(numbers...)
	println("Sum of numbers:", sumOfNumbersInSlice)
}

func sum(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
