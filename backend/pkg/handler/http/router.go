package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(
		gin.Logger(),
		gin.Recovery(),
	)

	router.GET("/articles", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "ok")
	})

	return router
}
