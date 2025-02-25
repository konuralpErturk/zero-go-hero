package main

import "fmt"

func main() {
	dog := Dog{
		Breed: "Labrador",
		Sound: "Woof",
	}
	dog.Bark()
	dog.Sound = "Bark!"
	dog.speakThreeTimes()
}

type Dog struct {
	Breed string
	Sound string
}

func (d Dog) Bark() {
	fmt.Printf("The %v dog barks %v\n", d.Breed, d.Sound)
}

func (d Dog) speakThreeTimes() {
	fmt.Printf("%v %v %v\n", d.Sound, d.Sound, d.Sound)
}