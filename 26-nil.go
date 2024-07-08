package main

import "fmt"

func nill() {
	// Nill
	// Nill adalah nilai default dari tipe data Map, Slice, Pointer, Channel, dan Function

	person := newMap("")
	if person == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(person)
	}
}

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}
