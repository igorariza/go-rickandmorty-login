package main

import (
	"fmt"

	"github.com/igorariza/Dockerized-Golang_API-MySql-React.js/api"
)

func main() {

	go api.MigrateExternalAPI()

	go func(msg string) {
		fmt.Println(msg)
	}("Starting migration external API")

	api.Run()

}
