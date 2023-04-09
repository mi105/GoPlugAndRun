package main

import (
	"awesomeProject/animal"
	"fmt"
)

type Dog struct{}

func (this Dog) Speak() {
	fmt.Println("Rough Rough!")
}

var ProcessDog animal.Animal = Dog{}
