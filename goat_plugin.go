package main

import (
	"awesomeProject/animal"
	"fmt"
)

type Goat struct{}

func (this Goat) Speak() {
	fmt.Println("Mehehehe Meheheh Mehehe!")
}

var ProcessGoat animal.Animal = Goat{}
