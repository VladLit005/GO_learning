package main

import (
	"errors"
	"fmt"
)

var ErrEmpty = errors.New("empty name")

func main() {
	fmt.Println(read(""))
	fmt.Println(read("test"))
}

func read(name string) error {
	if name == "" {
		return ErrEmpty
	}
	return nil
}
