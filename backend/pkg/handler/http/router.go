package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(
		gin.Logger(),
		gin.Recovery(),
		cors.New(cors.Config{
			AllowOrigins: []string{
				"http://localhost:5173",
			},
		}),
	)

	router.GET("/articles", func(ctx *gin.Context) {
		json := fmt.Sprintf(`
		{
			"message": "OK",
			"time": %v
		}`, time.Now())
		ctx.JSON(http.StatusOK, json)
	})

	return router
}
