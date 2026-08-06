package main

import (
	"fmt"
	"log"
)

func main() {
	divide(10, 2)              // 5
	divide(10, 0)              // Recover from panic: runtime error: integer divide by zero
	fmt.Println("after panic") // after panic
}

func divide(a, b int) {
	defer func() {
		if err := recover(); err != nil { // recover() возвращает панику которая случилась
			log.Println("Recover from panic:", err)
		}
	}()
	fmt.Println(a / b)
}
