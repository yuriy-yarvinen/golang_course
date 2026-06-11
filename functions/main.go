package main

func main() {

	numbers := []int{1, 2, 3, 4, 5}
	moreNumbers := []int{6, 7, 8, 9, 10}

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

	funcForNumbers := getTransformersFunction(&numbers)
	transformedNumbers = transformNumbers(&numbers, funcForNumbers)

	println("Transformed numbers using function from getTransformersFunction:", funcForNumbers)
	for _, number := range transformedNumbers {
		println(number)
	}

	funcForMoreNumbers := getTransformersFunction(&moreNumbers)
	transformedMoreNumbers := transformNumbers(&moreNumbers, funcForMoreNumbers)

	println("Transformed more numbers using function from getTransformersFunction:", funcForMoreNumbers)
	for _, number := range transformedMoreNumbers {
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

func getTransformersFunction(numbers *[]int) func(n int) int {
	if (*numbers)[0] == 1 {
		return double
	}
	return triple
}

func double(n int) int {
	return n * 2
}

func triple(n int) int {
	return n * 3
}
