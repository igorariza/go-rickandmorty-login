package api

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/igorariza/go-rickandmorty-login/api/controllers"
	"github.com/igorariza/go-rickandmorty-login/api/seed"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func init() {

	// loads values from .env into the system
	// if err := godotenv.Load(); err != nil {
	// 	log.Print("sad .env file found")
	// }
}

func Run() {

	//var err error
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println(" We are getting the env values")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	seed.Load(server.DB)
	port := os.Getenv("PORT")
	

	server.Run(":" + port)

}
func MigrateExternalAPI() {

	cmd, err := exec.Command("/bin/bash", "./entrypoint.sh").Output()
	if err != nil {
		fmt.Printf("error %s", err)
	}
	output := string(cmd)
	fmt.Println(output)

}
