package main

import (
	"fmt"
)

func main() {
	anInt := 42
	var p = &anInt
	fmt.Println("Value of p:", *p)
	fmt.Println("Address in p:", p)

	value1 := 42.13
	pointer1 := &value1
	fmt.Println("Value 1:", *pointer1)

	*pointer1 = *pointer1 / 32
	fmt.Println("Pointer 1:", *pointer1)
	fmt.Println("Value 1:", value1)
}
