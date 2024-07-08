package main

import "fmt"

// func main() {
// 	// Anonymous Function

// 	blacklist := func(name string) bool {
// 		return name == "admin"
// 	}

// 	register("admin", blacklist)
// 	register("root", blacklist)
// }

type BlackList func(string) bool

func register(name string, blackList BlackList) {
	if blackList(name) {
		fmt.Println("you are blocked ", name)
	} else {
		fmt.Println("Wellcome ", name)
	}
}