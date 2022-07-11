package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter value 1: ")
	val1, _ := reader.ReadString('\n')
	value1, err := strconv.ParseFloat(strings.TrimSpace(val1), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 1 must be a number")
	}

	fmt.Print("Enter value 2: ")
	val2, _ := reader.ReadString('\n')
	value2, err := strconv.ParseFloat(strings.TrimSpace(val2), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 2 must be a number")
	}

	sum := math.Round((value1+value2)*100) / 100
	fmt.Printf("The sum of %.2f and %.2f is %.2f", value1, value2, sum)

}
