package main

import (
	"github.com/btfidelis/quizzer/models"
	"github.com/gocraft/web"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

func main() {
	LoadConfig()
	models.BootDatabaseConn()

	router := web.New(Api{ResponseDefaultEncoding: "application/json"})
	SetMiddleware(router)
	SetRoutes(router)
	http.ListenAndServe("localhost:3000", router)
}
