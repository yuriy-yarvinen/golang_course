package main

import "fmt"

func main() {

	prices := []float64{19.99, 29.99, 9.99, 49.99}

	for i, price := range prices {
		fmt.Printf("Price %d: $%.2f\n", i+1, price)
	}

	fmt.Println("Slice:", prices[1:3])
	fmt.Println("Slice:", prices[1:])
	fmt.Println("Slice:", prices[:3])

	fmt.Println("Length:", len(prices))
	fmt.Println("Capacity:", cap(prices))

}
