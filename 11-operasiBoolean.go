package main

import "fmt"

func operasiboolean() {
	// operasi boolean => operasi yang menghasilkan nilai boolean
	// AND, OR, NOT

	// AND
	left := true
	right := false
	result := left && right
	fmt.Println(result) // false

	// OR
	left = true
	right = false
	result = left || right
	fmt.Println(result) // true

	// NOT
	left = true
	result = !left
	fmt.Println(result) // false

}
