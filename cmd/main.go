package main

import (
	"log"
	"net/http"
	"os"

	"todo-list/config"
	_ "todo-list/docs"
	"todo-list/router"

	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Todo List API
// @version 1.0
// @description This is a sample server for a todo list.
// @host localhost:8080
// @BasePath /
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.LoadConfig()
	r := router.SetupRouter()

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server started at http://localhost:" + port)
	log.Println("Documentation available at http://localhost:" + port + "/swagger/")

	log.Fatal(http.ListenAndServe(":"+port, r))
}
