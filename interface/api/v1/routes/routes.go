package routes

import (
	svcUser "pajarwp11/go-experiment/core/service/user"

	repoUser "pajarwp11/go-experiment/infrastructure/repository/mysql/user"

	handlerUser "pajarwp11/go-experiment/interface/api/user"

	"pajarwp11/go-experiment/config"

	"github.com/labstack/echo/v4"
)

func API(e *echo.Echo) {
	// instance DB
	db := config.MySQL

	// instace repository
	userRepo := repoUser.New(db)

	// instance service
	userSvc := svcUser.New(userRepo)

	// instance handler
	userHandler := handlerUser.New(userSvc)

	// register routes v1
	v1 := e.Group("/v1/api")

	gMerchant := v1.Group("/user")
	gMerchant.POST("", userHandler.RegisterUser)
	gMerchant.GET("/list", userHandler.GetUserList)

}
