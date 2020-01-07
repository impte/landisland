package main

import (
	"github.com/impte/landisland/api"
	"github.com/impte/landisland/config"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("配置文件加载错误")
	}
	log.Info("初始化系统配置")
	config.InitSetting()
	log.Info("初始化数据库")
	config.InitDatabase()
	log.Info("初始化redis")
	config.InitRedis()
	if err := api.Run(config.ServerPort); err != nil {
		panic(err.Error())
	}
}
