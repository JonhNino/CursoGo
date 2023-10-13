package main

import (
	"fmt"
)

func PrintList(list ...interface{}) {
	for _, value := range list {
		fmt.Println(value)
	}
}

type Numbers interface {
	~int | ~float64 | ~float32 | ~uint
}

func Sum[T Numbers](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

func Includes[T comparable](list []T, value T) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

func Filter[T comparable](list []T, callback func(T) bool) []T {
	result := make([]T, 0, len(list))
	for _, item := range list {
		if callback(item) {
			result = append(result, item)
		}
	}
	return result
}

type Prduct[T uint | string] struct {
	Id 		T
	Desc 	string
	Price 	float32
}

type integer int // defining the integer type

func main() {
	PrintList("Ariel", "Santiago", "Juan", "Pedro")
	PrintList(42, 12, 6, 3)
	PrintList(false, 12, 6, "Hola")
	fmt.Println(Sum(uint(3), uint(12), uint(6), uint(3)))
	fmt.Println(Sum(float32(34.4), float32(12), float32(6), float32(3)))

	var num1 integer = 100
	var num2 integer = 12341
	fmt.Println(Sum(num1, num2))

	strings := []string{"a", "b", "c", "d"}
	numbers := []int{1, 2, 3, 4,5,6,7}

	fmt.Println(Includes(strings, "a"))
	fmt.Println(Includes(strings, "b"))
	fmt.Println(Includes(strings, "f"))
	fmt.Println(Includes(numbers, 8))
	fmt.Println(Includes(numbers, 3))

	fmt.Println(Filter(numbers, func(value int) bool { return value < 4 }))
	fmt.Println(Filter(strings, func(value string) bool { return value > "b" }))

	product1 := Prduct[uint]{1,"Rey",145}
	product2 := Prduct[string]{"SDA-DASE23-DFASSG","Rey",145}
	fmt.Println(product1)
	fmt.Println(product2)
}
