package main

import (
	"fmt"
	"net/http"
	"sync"
)

type Cart map[string]int32

type ID int      // <- new type
type IDint = int // <- alias to type int

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

	nums := make([]int, 1, 3)
	fmt.Println("Nums:", nums)
	// [0]

	appendToSlice(nums, 1)
	fmt.Println("Nums after appendToSlice:", nums)
	// Nums after appendToSlice: [0]

	copySlice(nums, []int{2, 3})
	fmt.Println("Nums after copySlice:", nums)
	// Nums after copySlice: [2]

	// mutateSlice(nums, 1, 4)
	// fmt.Println("Nums after mutateSlice:", nums)
	// panic

	nums2 := []int{1, 2, 3}

	addNum(nums2[0:2])
	fmt.Println("Nums add:", nums2)

	addNums(nums2[0:2])
	fmt.Println("Nums addNums:", nums2)
	fmt.Println("Nums cap:", cap(nums2[0:2]))   // cap 3
	fmt.Println("Nums cap:", cap(nums2[0:2:2])) // cap 2

	// var c Cart // panic not inicialized
	// // var c Cart = make(Cart, 5)

	// c.addOrUpdate("orange", 2)
	// c.addOrUpdate("orange", 2)

	// fmt.Println(&c)

	newCart := make(Cart)
	newCart2 := &newCart

	newCart.addOrUpdate("ora", 2)
	(*newCart2).addOrUpdate("ora", 1)
	(*newCart2)["ora"] += 5

	fmt.Println(newCart)
	fmt.Println(*newCart2)

	number := 10
	funcAccum := accum()
	fmt.Println(funcAccum(number))
	fmt.Println(funcAccum(number))
	fmt.Println(funcAccum(number))
	fmt.Println(funcAccum(number))

	// type error interface
	// {*type=nil *data=nil}

	var err1 error           // {*type=nil *data=nil}
	fmt.Println(isNil(err1)) // true

	var err2 *errorCustom    // {*type=errorCustom *data=nil}
	fmt.Println(isNil(err2)) // false

	err2 = &errorCustom{msg: "error"}
	fmt.Println(isNil(err2)) // false

	err2 = nil               // обнулили только дату, а тип остался и всеравно ложно
	fmt.Println(isNil(err2)) // false

	err1 = err2
	fmt.Println(isNil(err1)) // false

	err1 = nil
	fmt.Println(isNil(err1))

	// urls := []string{"https://dzen.ru", "https://google.com", "https://ya.ru"}

	// wg := sync.WaitGroup{}
	// for _, url := range urls {
	// 	wg.Add(1)
	// 	go GetHttpData(url, &wg)
	// }

	// wg.Wait()

	set := []int{1, 2, 34, 1, 2}
	set2 := []int{1, 2, 3, 4, 5}

	fmt.Println(checkCoppies(set))
	fmt.Println(checkCoppies(set2))
}

func checkCoppies(slice []int) bool {
	checked := []int{}
	for _, number := range slice {
		if SliceContains(checked, number) {
			return true
		} else {
			checked = append(checked, number)
		}
	}
	return false
}

func SliceContains(slice []int, number int) bool {
	for _, numberInSlice := range slice {
		if number == numberInSlice {
			return true
		}
	}
	return false
}

func GetHttpData(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	response, errHttp := http.Get(url)
	if errHttp != nil {
		fmt.Println(errHttp)
		return
	}
	fmt.Println(response)
}

type errorCustom struct {
	msg string
}

func (e errorCustom) Error() string {
	return e.msg
}

func isNil(err error) bool {
	return err == nil
}

func (cart Cart) addOrUpdate(sku string, count int32) {
	cart[sku] += count
}

func appendToSlice(slice []int, value int) {
	slice = append(slice, value)
}

func copySlice(dst []int, src []int) {
	copy(dst, src)
}

func mutateSlice(slice []int, index int, newValue int) {
	slice[index] = newValue
}

func addNum(nums []int) {
	nums = append(nums, 4)
}

func addNums(nums []int) {
	nums = append(nums, 5, 6)
}

func accum() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
