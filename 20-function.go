package main

import "fmt"

func sayHello(name string) {
	fmt.Println("Hello", name)
}

// function with return value
func getHello(name string) string {
	return "Hello " + name
}

// function with multiple return value
func getHello2() (string, string) {
	return "fairuz", "ulum"
}

// Named return value
func getHello3() (firstName, midleName, lastName string) {
	firstName = "ahmad"
	midleName = "fairuz"
	lastName = "ulum"
	return firstName, midleName, lastName

}
// func main() {
// 	sayHello("Fairuz")
// 	fmt.Println(getHello("Fairuz"))
// 	firstName, lastName := getHello2()
// 	fmt.Println(firstName, lastName)

// 	// jika kita hanya ingin mengambil salah satu return value
// 	namaDepan, _ := getHello2()
// 	fmt.Println(namaDepan)

// 	// Named return value
// 	_, midleName, _ := getHello3()
// 	fmt.Println(midleName)

// }

