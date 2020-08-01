package main

import (
	"fmt"
	"heokjin/kingofweather-api/model"
	"log"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	model.InitDB()
	//model.InitSchema()

	fmt.Println("TEST1")

	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//e.Use(middleware.Static("static"))

	e.GET("/checkout", CheckOut)

	// Start server
	e.Start(":" + port)
}
