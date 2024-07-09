package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Arguments: ", args)

	fmt.Println("-: ", args[1])
	fmt.Println("-: ", args[2])

	// mengambil detail nama hostname
	name, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("hostName: ", name)
	}

	// mengambil env
	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")
	fmt.Println("username: ", username)
	fmt.Println("password: ", password)

}
