package config

import (
	"errors"
	"github.com/go-redis/redis"
	"github.com/impte/landisland/utils"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
	"time"
)

func InitSetting() {
	ServerPort = os.Getenv("server.port")
	DataSourceUrl = os.Getenv("datasource.url")
	TokenSecret = os.Getenv("landisland.token.secret")
	RedisUrl = os.Getenv("redis.url")
	RedisPassword = os.Getenv("redis.password")
}

func InitDatabase() {
	if utils.StringIsEmpty(DataSourceUrl) {
		panic(errors.New("请先初始化数据库地址"))
	}
	db, err := gorm.Open("mysql", DataSourceUrl)
	if err != nil {
		panic(err.Error())
	}
	Database = db
	log.Info("数据库连接成功")
}

func InitRedis() {
	if utils.StringIsEmpty(RedisUrl) || utils.StringIsEmpty(RedisPassword) {
		panic(errors.New("请先初始化redis地址与密码"))
	}
	expire, err := strconv.ParseInt(os.Getenv("redis.expire"), 10, 64)
	if err != nil {
		panic(errors.New("缓存过期时间类型错误"))
	}
	RedisExpire = time.Duration(expire)
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     RedisUrl,
		Password: RedisPassword, // no password set
		DB:       0,  // use default DB
	})
	_, clientErr := RedisClient.Ping().Result()
	if clientErr != nil {
		panic(clientErr)
	}
}