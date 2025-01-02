package main

//Directory: /main.go
//This file is the application entry file

import (
	"log"

	"github.com/fauzancodes/crud-echo/app/config"
	"github.com/fauzancodes/crud-echo/app/routes"
	"github.com/labstack/echo/v4"

	//This package below is used to load environment variables automatically every time the application runs
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	//App initiation
	app := Init()

	//Get port from environment variable
	port := config.LoadConfig().Port

	log.Println("Server running on port:" + port)

	//Run the application on the specified port
	app.Logger.Fatal(app.Start(":" + port))
}

func Init() *echo.Echo {
	//Echo initiation
	app := echo.New()

	//Database initiation
	config.Database()

	//Routing initiation
	routes.Route(app)

	return app
}
