package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)
	fmt.Printf("%T\n", colors)
	colors = append(colors, "Purple")
	fmt.Println(colors)

	colors = append(colors[1:])
	fmt.Println(colors)

	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	numbers := make([]int, 5, 5)
	numbers[0] = 134
	numbers[1] = 83
	numbers[2] = 85
	numbers[3] = 62
	numbers[4] = 79
	fmt.Println(numbers)

	numbers = append(numbers, 235)
	fmt.Println(numbers)

	sort.Ints(numbers)
	fmt.Println(numbers)
}
