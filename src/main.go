package src

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/luiz-vinholi/vmy-users-crud/src/app"
	"github.com/luiz-vinholi/vmy-users-crud/src/infra/databases"
)

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error on loading .env file.")
	}
	databases.ConnectMongoDB()
	app.Load()
}
