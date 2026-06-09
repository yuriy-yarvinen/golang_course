package main

func main() {
	println(plus(3, 5))                // Output: 8
	println(plus(3.5, 2.5))            // Output: 6.0
	println(plus("Hello, ", "world!")) // Output: Hello, world!
}

func plus[T int | float64 | string](a, b T) T {
	return a + b
}
