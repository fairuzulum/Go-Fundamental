package database

import "fmt"

var connection string

func init() {
	connection = "MySql"
}

func GetDatabase() {
	fmt.Println(connection)
}
