package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// const aConst int = 64

func main() {
	// Initialize Variables
	/*
	var aString string = "This is Go!"

	fmt.Println(aString)
	fmt.Printf("The variable's type is %T\n", aString)

	var anInteger int = 42
	fmt.Println(anInteger)
	fmt.Printf("The variable's type is %T\n", anInteger)

	var defaultInt int
	fmt.Println(defaultInt)
	fmt.Printf("The variable's type is %T\n", defaultInt)

	var anotherString = "This is another string"
	fmt.Println(anotherString)
	fmt.Printf("The variable's type is %T\n", anotherString)

	var anotherInt = 53
	fmt.Println(anotherInt)
	fmt.Printf("The variable's type is %T\n", anotherInt)

	myString := "This is also a string"
	fmt.Println(myString)
	fmt.Printf("The variable's type is %T\n", myString)

	fmt.Println(aConst)
	fmt.Printf("The variable's type is %T\n", aConst)
	*/

	// Read from Command Line
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered:", input)

	fmt.Print("Enter a number: ")
	numInput, _ := reader.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of number:", aFloat)
	}


}