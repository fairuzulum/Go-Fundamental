package main

import "fmt"

func switchExpression() {
	// switch expression
	name := "ulum"

	switch name {
	case "fairuz":
		fmt.Println("Hello fairuz")
	case "ulum":
		fmt.Println("Hello ulum")
	default:
		fmt.Println("Hello, boleh kenalan?")
	}

	// switch dengan short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// switch tanpa kondisi
	length := len(name)
	switch {
	case length > 5:
		fmt.Println("Nama terlalu panjang")
	case length < 5:
		fmt.Println("Nama terlalu pendek")
	}

}
