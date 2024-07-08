package main

import "fmt"

// func main() {
// 	// Function as Parameter
// 	// function juga bisa dijadikan sebagai parameter untuk function lain

// 	sayHelloWithFilter("Anjing", spamFilter)

// 	filter := spamFilter
// 	sayHelloWithFilter("Fairuz", filter)


// }


func sayHelloWithFilter(name string, filter func(string) string) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}