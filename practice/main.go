package main

import "fmt"

type Product struct {
	title string
	id    int
	price float64
}

func main() {

	hobbies := [3]string{"Reading", "Cooking", "Traveling"}

	fmt.Println("Hobbies:", hobbies[0], hobbies[1], hobbies[2])
	fmt.Println("First hobby:", hobbies[0])
	fmt.Println("Second and third hobbies:", hobbies[1], hobbies[2])

	slice1 := hobbies[0:2]
	slice2 := hobbies[:2]
	fmt.Println("Slice 1:", slice1)
	fmt.Println("Slice 2:", slice2)

	slice1 = hobbies[1:3]
	fmt.Println("Resliced Slice 1:", slice1)

	courseGoals := []string{"Learn Go", "Build a project"}
	fmt.Println("Course Goals:", courseGoals[0], courseGoals[1])

	courseGoals[1] = "Master Go"
	fmt.Println("Updated Course Goals:", courseGoals[0], courseGoals[1])

	courseGoals = append(courseGoals, "Master Go Concurrency")
	fmt.Println("Final Course Goals:", courseGoals[0], courseGoals[1], courseGoals[2])

	products := []Product{
		{title: "Laptop", id: 1, price: 999.99},
		{title: "Smartphone", id: 2, price: 499.99},
	}

	fmt.Println("Products:")
	for _, product := range products {
		fmt.Println("Title:", product.title, "ID:", product.id, "Price:", product.price)
	}

	products = append(products, Product{title: "Tablet", id: 3, price: 299.99})
	fmt.Println("Updated Products:")
	for _, product := range products {
		fmt.Println("Title:", product.title, "ID:", product.id, "Price:", product.price)
	}
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
