package main

import "fmt"

type Customer struct {
	Name string
}

func (c Customer) sayHi() {
	fmt.Println("Hello", c.Name)
}

// func main() {
// 	fairuz := Customer{"Fairuz"}
// 	fairuz.sayHi()
// }