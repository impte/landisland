package main

import (
	"github.com/impte/landisland/api"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("配置文件加载错误")
	}
	server := api.Api{
		Listen: os.Getenv("listen"),
		DbUrl: os.Getenv("DbUrl"),
	}
	if err := server.Run(); err != nil {
		panic(err.Error())
	}
}
