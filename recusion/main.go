package main

import (
	"fmt"
	"math/big"
)

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127, 128, 129, 130, 131, 132, 133, 134, 135, 136, 137, 138, 139, 140, 141, 142, 143, 144, 145, 146, 147, 148, 149, 150}

	println("Original numbers:")
	for _, number := range numbers {
		fmt.Println(number, " = ", factorial(number))
	}

}

// / int is not enough to store the result of factorial for numbers greater than 20 its
//
// 51 090 942 171 709 440 000
// int 64 bit is -  9 223 372 036 854 775 807
//
// so we need to use big.Int from math/big package to handle large integers.
func factorial(number int) *big.Int {
	if number <= 0 {
		return big.NewInt(1)
	}
	return new(big.Int).Mul(big.NewInt(int64(number)), factorial(number-1))
}
