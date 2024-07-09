package main

import "fmt"

// Struct adalah kumpulan definisi variabel (atau property) dan fungsi (atau method) yang dibungkus sebagai sebuah tipe data baru.

type Person struct {
	Name string
	Age  int
	Address string
}

func mainStruck() {

	fairuz := Person{
		Name: "Fairuz",
		Age:  22,
		Address: "Jakarta",
	}
	fmt.Println(fairuz)

	ulum := Person{"Ulum", 20, "Jakarta"}
	fmt.Println(ulum)

	
}