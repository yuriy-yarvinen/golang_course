package main

import "fmt"

type Number interface {
	int | float64
}

func sumMapValues[T Number](m map[string]T) T {
	var sum T
	for _, val := range m {
		sum += val
	}
	return sum
}

type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(elem T) {
	s.data = append(s.data, elem)
}

func (s *Stack[T]) Pop() T {
	if len(s.data) == 0 {
		// важно нужно объявить
		var elem T
		return elem
	}
	pos := len(s.data) - 1
	elem := s.data[pos]
	s.data = s.data[:pos]
	return elem
}

type NumberStruct struct {
	numf float64
	num  int
}

// comparable - значит на типе определены операции ==, !=

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](elems ...T) Set[T] {
	set := make(map[T]struct{}, len(elems))
	for _, elem := range elems {
		set[elem] = struct{}{}
	}
	return set
}

func (s Set[T]) Put(elem T) {
	s[elem] = struct{}{}
}

func (s Set[T]) Has(elem T) bool {
	_, has := s[elem]
	return has
}

type strictInt interface {
	int8 | int16 | int32 | int64 | int
}

type flexibleInt interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int
}

type myCustomInt int

func strictSum[T strictInt](a, b T) T {
	return a + b
}

func flexibleSum[T flexibleInt](a, b T) T {
	return a + b
}

func main() {

	var a, b int
	strictSum(a, b)
	flexibleSum(a, b)

	var c, d myCustomInt
	// strictSum(c, d) // ошибка компиляции
	flexibleSum(c, d)

	println(plus(3, 5))                // Output: 8
	println(plus(3.5, 2.5))            // Output: 6.0
	println(plus("Hello, ", "world!")) // Output: Hello, world!

	map1 := map[string]int{
		"one": 1,
		"two": 2,
	}
	map2 := map[string]float64{
		"one": 1.2,
		"two": 2.5,
	}
	fmt.Println(sumMapValues(map1))
	fmt.Println(sumMapValues(map2))

	intStack := NewStack[int]()
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)
	fmt.Println(intStack.Pop())
	fmt.Println(intStack.Pop())
	fmt.Println(intStack.Pop())
	fmt.Println(intStack.Pop())
	fmt.Println(intStack.Pop())

	strStack := NewStack[string]()
	strStack.Push("sdf")
	strStack.Push("sdfdfs")
	fmt.Println(strStack.Pop())
	fmt.Println(strStack.Pop())
	fmt.Println(strStack.Pop())

	nsStack := NewStack[NumberStruct]()
	nsStack.Push(NumberStruct{numf: 0.2, num: 1})
	nsStack.Push(NumberStruct{numf: 1.2, num: 2})
	fmt.Println(nsStack.Pop())
	fmt.Println(nsStack.Pop())
	fmt.Println(nsStack.Pop())

	set := NewSet("sdf", "dfsdf", "dsfsdf")
	set.Put("dddd")

	fmt.Println("sdf", set.Has("sdf"))
	fmt.Println("sdddf", set.Has("sdddf"))
}

func plus[T int | float64 | string](a, b T) T {
	return a + b
}
