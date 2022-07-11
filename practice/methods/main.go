package main

import "fmt"

type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

func (d Dog) Bark() {
	fmt.Println(d.Sound)
}

func (d Dog) BarkThreeTimes() {
	d.Sound = fmt.Sprintf("%v %v %v", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}

func main() {
	poodle := Dog{"Poodle", 10, "Woof!"}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)

	poodle.Bark()
	poodle.Sound = "Arf!"
	poodle.Bark()
	poodle.BarkThreeTimes()
	poodle.Bark()
}
