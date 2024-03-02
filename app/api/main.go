package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"pajarwp11/go-experiment/config"
	"pajarwp11/go-experiment/interface/api/v1/routes"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.elastic.co/apm/module/apmechov4/v2"
)

func main() {
	config.LoadEnvVars()
	config.MysqlConnect()
	config.ConnectMongoDB()
	AppStart()
}

func AppStart() {
	e := echo.New()

	e.Use(apmechov4.Middleware())
	e.Use(middleware.Recover())

	routes.API(e)

	go func() {
		if err := e.Start(":" + os.Getenv("APP_PORT")); err != nil {
			log.Fatal("Shutting down the server.")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err.Error())
	}
}
