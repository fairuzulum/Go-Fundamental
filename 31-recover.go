package main

import "fmt"

func log() {
	fmt.Println("Logging Done")
	message := recover()
	if message != nil {
		fmt.Println("Error Message:", message)
	}
}

func runApp(error bool) {
	defer log()
	if error{
		panic("APLIKASI ERROR")
	}

	fmt.Println("APLIKASI BERJALAN")


}

func mainRecover() {
	runApp(false)

}