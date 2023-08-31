package api

import (
	//"fmt"

	"github.com/igorariza/go-rickandmorty-login/api"
)

func main() {

	go api.MigrateExternalAPI()

	// go func(msg string) {
	// 	fmt.Println(msg)
	// }("Starting migration external API")

	api.Run()

}
