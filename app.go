package main

import (
	"github.com/kiesproject/prepass-api/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 2,
	}))
	e.Use(middleware.Secure())

	//CORS
	e.Use(middleware.CORS())

	// Prepassエンドポイントグループ
	prepass := e.Group("/prepass/:version")

	// Routers
	prepass.GET("/search", handler.GetSearch)

	e.Start(":8080")
}
