package main

func constant() {
	// Constant
	// Constant adalah tempat untuk menyimpan data
	// Constant bersifat immutable, artinya nilainya tidak dapat diubah setelah dideklarasikan
	// Constant tidak bisa di deklarasikan tanpa value
	// Constant bisa di deklarasikan di package manapun
	// Constant bersifat static type, artinya setiap constant harus memiliki tipe data yang sama
	// Constant yang tidak digunakan tidak akan menyebabkan error

	// Deklarasi Constant
	const name string = "Fairuz"
	println(name)

	// Deklarasi Constant dengan multiple value
	const (
		firstName = "Fairuz"
		lastName  = "Ulum"
	)
	println(firstName, lastName)
}