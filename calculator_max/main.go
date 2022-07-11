package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getInputFromConsole(inputMsg string) string {
	fmt.Printf(inputMsg)
	val, _ := reader.ReadString('\n')
	return strings.TrimSpace(val)
}

func getNumInputFromConsole(valNum string) float64 {
	val := getInputFromConsole(fmt.Sprintf("Value %s: ", valNum))
	value, err := strconv.ParseFloat(strings.TrimSpace(val), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value " + valNum + " must be a number")
	}

	return value
}

func roundValue(value float64) float64 {
	return math.Round(value*100) / 100
}

func addValues(value1 float64, value2 float64) float64 {
	return roundValue(value1 + value2)
}
func subtractValues(value1 float64, value2 float64) float64 {
	return roundValue(value1 - value2)
}
func multiplyValues(value1 float64, value2 float64) float64 {
	return roundValue(value1 * value2)
}
func divideValues(value1 float64, value2 float64) float64 {
	return roundValue(value1 / value2)
}

func main() {

	value1 := getNumInputFromConsole("1")
	value2 := getNumInputFromConsole("2")
	op := getInputFromConsole("Select an operation (+ - * /): ")

	var result float64
	switch op {
	case "+":
		result = addValues(value1, value2)
	case "-":
		result = subtractValues(value1, value2)
	case "*":
		result = multiplyValues(value1, value2)
	case "/":
		result = divideValues(value1, value2)
	default:
		panic("Operation is not from the available options (+ - * /)")
	}

	fmt.Printf("The result is %.2f", result)
}
