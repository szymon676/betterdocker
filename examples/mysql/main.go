package main

import (
	"log"

	"github.com/szymon676/betterdocker/mysql"
)

func main() {
	// Define options. Leave empty struct and they will automaticly fill with default settings
	opts := &mysql.MySQLContainerOptions{}

	// Initalize container struct
	container := mysql.NewMySQLContainer(opts)

	// Run container
	err := container.Run()
	if err != nil {
		log.Fatal(err)
	}
	//stop container
	defer container.Stop()
}
