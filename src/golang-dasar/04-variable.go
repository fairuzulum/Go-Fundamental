package main

func variable() {
	// Variable
	// Variable adalah tempat untuk menyimpan data
	// Variable harus dideklarasikan terlebih dahulu sebelum digunakan
	// Variable di Go bersifat static type, artinya setiap variable harus memiliki tipe data yang sama
	// Variable yang tidak digunakan akan menyebabkan error

	// Deklarasi Variable
	var name string
	name = "Fairuz"
	println(name)

	// Deklarasi Variable dengan tipe data
	var age int
	age = 20
	println(age)

	// Deklarasi Variable dengan tipe data dan value
	var address string = "Jakarta"
	println(address)

	// Deklarasi Variable tanpa tipe data
	country := "Indonesia"
	println(country)

	// Deklarasi Variable dengan multiple value
	var (
		firstName = "Fairuz"
		lastName  = "Ulum"
	)
	println(firstName, lastName)
}
