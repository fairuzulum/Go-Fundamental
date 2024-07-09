package main

import "fmt"

type HashName interface {
	getName() string
}

func sayAnimal(hashName HashName) {
	fmt.Println(hashName.getName())
}

type Animal struct {
	name string
}



func (animal Animal) getName() string {
	return animal.name
}

// func main() {
// 	var anjing Animal
// 	anjing.name = "Anjing"
// 	sayAnimal(anjing)

// }