package main

import (
	"fmt"
	"time"
)

func main() {
	n := time.Now()
	fmt.Println("I typed this code at", n)

	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("Go launched at", t)
	fmt.Println(t.Format(time.ANSIC))

	parsedTime, err := time.Parse(time.ANSIC, "Tue Nov 10 23:00:00 2009")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The type of parsedTime is %T\n", parsedTime)
	}
}
