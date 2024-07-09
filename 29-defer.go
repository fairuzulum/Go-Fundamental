package main

import "fmt"

func logging() {
	fmt.Println("Logging Done")

}

func runApplication() {
	defer logging()
	fmt.Println("Run Application")
}

func mainDefer() {
	// defer => menjalankan fungsi setelah fungsi yang bersangkutan selesai dijalankan
	runApplication()
}
