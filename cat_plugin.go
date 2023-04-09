package main

import (
	"awesomeProject/animal"
	"fmt"
)

type Cat struct{}

func (this Cat) Speak() {
	fmt.Println("Meow meow!")
}

var ProcessCat animal.Animal = Cat{}
