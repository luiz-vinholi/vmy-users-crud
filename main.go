package main

import (
	"log"
	"vmytest/src/app"
	"vmytest/src/infra/databases"
	"vmytest/src/interfaces/rest"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error on loading .env file.")
	}
	databases.ConnectMongoDB()
	app.Load()
	rest.Run()
}
