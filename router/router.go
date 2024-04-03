package router

import (
	"free-gpt3.5-2api/middleware"
	"free-gpt3.5-2api/v1/chat"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRouter(router *gin.Engine) {

	router.GET("/", Index)
	v1Router := router.Group("/v1")
	v1Router.Use(middleware.V1Cors)
	v1Router.Use(middleware.V1Request)
	//v1Router.Use(middleware.V1Auth)
	v1Router.Use(middleware.V1Response)
	v1Router.POST("/chat/completions", chat.Completions)
}

func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello,This is free-gpt3.5-2api.")
}
