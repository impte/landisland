package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/impte/landisland/auth"
	"github.com/impte/landisland/controller"
)

func Run(address string) error {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "token", "Content-Type"},
		AllowCredentials: true,
	}))
	router.Use(auth.TokenVerify())
	router.POST("/api/login", controller.Login)
	router.POST("/api/logout", controller.Logout)
	return router.Run(address)
}
