package main

import "fmt"

// func main() {
// 	a := []int{1, 2, 3}
// 	b := a[5]
// 	fmt.Println(b) // panic: runtime error: index out of range
// }

// func main() {
// 	panic("something goes wrong!")
// }

func main() {
	defer func() {
		fmt.Println("ОК")
	}()
	panic("something goes wrong!")
	/*
		ОК
		panic: something goes wrong!
	*/ // defer раньше panic
}
