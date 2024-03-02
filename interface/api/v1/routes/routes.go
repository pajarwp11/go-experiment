package routes

import (
	svcItem "pajarwp11/go-experiment/core/service/item"
	svcUser "pajarwp11/go-experiment/core/service/user"

	repoItem "pajarwp11/go-experiment/infrastructure/repository/mongo/item"
	repoUser "pajarwp11/go-experiment/infrastructure/repository/mysql/user"

	handlerItem "pajarwp11/go-experiment/interface/api/item"
	handlerUser "pajarwp11/go-experiment/interface/api/user"

	"pajarwp11/go-experiment/config"

	"github.com/labstack/echo/v4"
)

func API(e *echo.Echo) {
	// instance DB
	db := config.MySQL
	mongo := config.MongoDB

	// instace repository
	userRepo := repoUser.New(db)
	itemRepo := repoItem.New(mongo)

	// instance service
	userSvc := svcUser.New(userRepo)
	itemSvc := svcItem.New(itemRepo)

	// instance handler
	userHandler := handlerUser.New(userSvc)
	itemHandler := handlerItem.New(itemSvc)

	// register routes v1
	v1 := e.Group("/v1/api")

	gUser := v1.Group("/user")
	gUser.POST("", userHandler.RegisterUser)
	gUser.GET("/list", userHandler.GetUserList)
	gUser.GET("/:id", userHandler.GetUserData)
	gUser.PUT("/:id", userHandler.UpdateUser)
	gUser.DELETE("/:id", userHandler.DeleteUser)

	gItem := v1.Group("/item")
	gItem.POST("", itemHandler.InsertItem)
	gItem.GET("/:id", itemHandler.GetItemData)
	gItem.PUT("/:id", itemHandler.UpdateItem)
	gItem.DELETE("/:id", itemHandler.DeleteItem)
}
