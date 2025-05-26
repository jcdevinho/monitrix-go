package main

import (
	"monitrix-api/routes"
	"monitrix-api/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	utils.InitDB()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":8085"))//escolhi porque minha vps esta usando do 8080:8084 lol
}
