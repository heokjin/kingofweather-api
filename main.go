package main

import (
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

	//TODO 배포하기전에 주석을 풀자!
	model.InitDB()

	go GoScheduleDeleteDB()
	go GoScheduleMidLandFcst()
	go GoScheduleMidTemp()


	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//e.Use(middleware.Static("static"))

	e.GET("/check", Check)

	e.GET("/getTmFc", GetWeatherMidWater)

	// Start server
	e.Start(":" + port)
}
