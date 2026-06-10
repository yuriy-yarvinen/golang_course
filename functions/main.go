package main

func main() {

	numbers := []int{1, 2, 3, 4, 5}

	println("Original numbers:")
	for _, number := range numbers {
		println(number)
	}

	transformedNumbers := transformNumbers(&numbers, double)

	println("Doubled numbers:")
	for _, number := range transformedNumbers {
		println(number)
	}

	transformedNumbers = transformNumbers(&numbers, triple)

	println("Tripled numbers:")
	for _, number := range transformedNumbers {
		println(number)
	}

}

func transformNumbers(numbers *[]int, multiplierFunc func(int) int) []int {
	transformed := make([]int, len(*numbers))

	for index, number := range *numbers {
		transformed[index] = multiplierFunc(number)
	}

	return transformed
}

func double(n int) int {
	return n * 2
}

func triple(n int) int {
	return n * 3
}
