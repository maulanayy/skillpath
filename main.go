package main

import (
	"context"
	"fmt"
	"log"
	"skillpath/app"
	"skillpath/handler"
	"skillpath/model"
	"skillpath/pkg"
	"skillpath/route"
	"skillpath/usecase"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tylerb/graceful"
)

func main() {

	filePath := "./config.yaml"

	var cfg model.Config
	err := pkg.ReadFile(&cfg, filePath)
	if err != nil {
		log.Fatal(err)
	}

	db, err := pkg.NewDatabase(context.Background(), model.DBParams{
		Address:  cfg.DB.Address,
		DBName:   cfg.DB.Name,
		Port:     cfg.DB.Port,
		SSLMode:  cfg.DB.SSL,
		User:     cfg.DB.User,
		Password: cfg.DB.Password,
	})
	if err != nil {
		log.Fatal(err)
	}

	repoDB := app.NewRepo(db)
	coeUC := usecase.NewUsecase(repoDB, &cfg)
	handler := handler.NewHandler(coeUC, &cfg)

	e := echo.New()

	route.RegisterRoute(e, handler)
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	e.Server.Addr = ":" + cfg.Rest.Port
	fmt.Println("SERVER RUNNING IN PORT  : ", cfg.Rest.Port)
	graceful.ListenAndServe(e.Server, 10*time.Second)

}
