package main

import "fmt"

func logg() {
	fmt.Println("Logging Done")
}

func applicationRun(error bool) {
	defer logg()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("APLIKASI BERJALAN")
}

func mainPanic() {
	// Panic
	// Panic adalah fungsi yang digunakan untuk menghentikan program
	// Panic biasanya dipanggil ketika terjadi kesalahan pada program
	// Panic akan menghentikan program dan menampilkan pesan error
	applicationRun(true)

}
