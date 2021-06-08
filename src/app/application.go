package app

import (
	"github.com/juliandev/bookstore_oauth-api/src/services/access_token"
	"github.com/juliandev/bookstore_oauth-api/src/repository/db"
	"github.com/juliandev/bookstore_oauth-api/src/repository/rest"
	"github.com/juliandev/bookstore_oauth-api/src/http"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	dbRepository := db.NewRepository()
	usersRepository := rest.NewRestUsersRepository()
	atService := access_token.NewService(dbRepository,usersRepository)
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	router.PATCH("/oauth/access_token", atHandler.Update)
	router.Run(":8090")
}
