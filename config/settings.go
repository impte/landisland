package config

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"time"
)

var (
	ServerPort    string
	DataSourceUrl string
	TokenSecret   string
	Database      *gorm.DB

	RedisUrl      string
	RedisPassword string
	RedisExpire   time.Duration
	RedisClient   *redis.Client
)
