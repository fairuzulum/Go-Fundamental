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

}
