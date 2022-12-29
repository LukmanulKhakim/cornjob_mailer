package main

import (
	"cornjobmailer/config"
	"cornjobmailer/features/user/delivery"
	"cornjobmailer/features/user/repository"
	"cornjobmailer/features/user/services"
	"cornjobmailer/utils/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.NewConfig()
	db := database.InitDB(cfg)

	uRepo := repository.New(db)
	uService := services.New(uRepo)
	delivery.New(e, uService)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.Logger.Fatal(e.Start(":8000"))
}
