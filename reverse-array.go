package main

import "fmt"

func main() {
	var length int
	fmt.Print("Enter the size of the array: ")
	fmt.Scanln(&length)
	// var a []int
	a := make([]int, length)
	b := make([]int, length)
	// b := make([]int, length)
	println("Enter the array:")
	for i := 0; i < length; i++ {
		fmt.Scan(&a[i])
	}
	b = reverse(a)
	for i := 0; i < length; i++ {
		fmt.Print(b[i], " ")
	}
}

func reverse(a []int) []int {
	var count int
	count = len(a) - 1
	b := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		b[count] = a[i]
		count--
	}
	return b
}
