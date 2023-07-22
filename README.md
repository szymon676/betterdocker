# library that allows you to run easily docker containers for mysql, postgres or mongo

## MySQL

```go 
package main

import "github.com/szymon676/betterdocker/mysql"

func main() {
	// Define options. Leave empty struct and they will automaticly fill with default settings
	opts := &mysql.MySQLContainerOptions{}

	// Initialize container struct
	container := mysql.NewMySQLContainer(opts)

	// Run container
	err := container.Run()
	if err != nil {
		log.Fatal(err)
	}
	//stop container
	defer container.Stop()
}
```

## PostgreSQL

```go
package main

import (
	"log"

	"github.com/szymon676/betterdocker/postgres"
)

func main() {
	// Define options. Leave empty struct and they will automaticly fill with default settings
	opts := &postgres.PostgresContainerOptions{}

	// Initalize container struct
	container := postgres.NewPostgresContainer(opts)

	// run container
	err := container.Run()
	if err != nil {
		log.Fatal(err)
	}

	// stop container
	defer container.Stop()
}
```