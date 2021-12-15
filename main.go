package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4/server"
	"send-email/controller"
	"send-email/libs/oauth"
	"send-email/model/repository"
)

func main() {
	repository.Init()
	// Initialize the oauth service
	oauth.InitServer()
	oauth.SetAllowGetAccessRequest(true)
	oauth.SetClientInfoHandler(server.ClientFormHandler)

	router := gin.Default()
	auth := router.Group("/oauth2")
	{
		auth.GET("/token", oauth.HandleTokenRequest)
	}
	api := router.Group("/api")
	{
		api.Use(oauth.HandleTokenVerify())
		api.GET("/logs", controller.GetEmailLogs)
		api.GET("/logs/:id", controller.GetEmailLogById)
	}

	err := router.Run("localhost:8030")
	if err != nil {
		return
	}
}
