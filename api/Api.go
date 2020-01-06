package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/impte/landisland/auth"
	"github.com/impte/landisland/constant"
	"github.com/impte/landisland/controller"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type Api struct {
	Listen string
	DbUrl string
}

func (a *Api) Run() error {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "token", "Content-Type"},
		AllowCredentials: true,
	}))
	initDatabase(a.DbUrl)
	router.Use(auth.TokenVerify())
	router.POST("/api/login", controller.Login)
	router.POST("/api/logout", controller.Logout)
	return router.Run(a.Listen)
}

func initDatabase(dbUrl string) {
	db, err := gorm.Open("mysql", dbUrl)
	if err != nil {
		panic(err.Error())
	}
	constant.Database = db
	log.Info("数据库连接成功")
}
