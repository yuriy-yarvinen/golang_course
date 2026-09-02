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

	newNumbers := []int{11, 12, 13, 14, 15}

	transformedNewNumbers := transformNumbers(&newNumbers, func(n int) int {
		return n * 4
	})
	println("Transformed new numbers using anonymous function:")

	for _, number := range transformedNewNumbers {
		println(number)
	}

	multiplierByfiveFunction := createTrasformerFunction(5)
	transformedNewNumbers = transformNumbers(&newNumbers, multiplierByfiveFunction)

	println("Transformed new numbers using function from createTrasformerFunction with multiplier 5:")
	for _, number := range transformedNewNumbers {
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

func createTrasformerFunction(multiplier int) func(number int) int {
	return func(number int) int {
		return number * multiplier
	}
}
