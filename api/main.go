package api

import (
	//"fmt"

	//"github.com/igorariza/go-rickandmorty-login/api"
)

func main() {

	go MigrateExternalAPI()

	// go func(msg string) {
	// 	fmt.Println(msg)
	// }("Starting migration external API")

	Run()

}
