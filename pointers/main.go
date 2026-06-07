package main

import "fmt"

func main() {
	age := 36

	agePointer := &age
	fmt.Println("Age pointer:", agePointer)
	fmt.Println("Age value through pointer:", *agePointer)
	*agePointer = 37

	fmt.Println("Age:", age)
	adultYears := getAdultYears(agePointer)
	fmt.Println("Adult years:", adultYears)
}

func getAdultYears(age *int) int {
	return *age - 18
}
