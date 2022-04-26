package main

import (
	"fmt"
	"log"
	"os"
	"url-shortener/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

//putting all routes in the below function
func setupRoutes(app *fiber.App){
	app.Get("/:url", routes.ResolveUrl)
	app.Post("/api/v1", routes.ShortenUrl)
}

func main(){
	err := godotenv.Load()

	if(err!=nil){
		fmt.Println(err)
	}
	//app above works same as how it works in express
	app := fiber.New()

	//adding logger to keep logs in place
	app.Use(logger.New())

	//calling routes function
	setupRoutes(app)

	//starting the server, and taking port from the .env file 
	//fatal is like print in log followed by os.Exit(1), basically exiting 
	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}