package main

import (
	app "gocrud/App"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	myApp := app.App{}
	myApp.Initialise()
	myApp.Run(os.Getenv("URL"))
	myApp.HandleRoutes()
}
