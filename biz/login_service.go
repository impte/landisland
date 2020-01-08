package biz

import (
	"encoding/json"
	"errors"
	"github.com/impte/landisland/common/cachekey"
	"github.com/impte/landisland/config"
	"github.com/impte/landisland/dao"
	"github.com/impte/landisland/model"
	"github.com/impte/landisland/reqres"
	"github.com/impte/landisland/utils"
	log "github.com/sirupsen/logrus"
	"time"
)

func Login(request reqres.LoginRequest) (reqres.LoginResponse, error) {
	enPassword, err := utils.Md5(request.Password)
	if err != nil {
		return reqres.LoginResponse{}, err
	}
	account, err := getAccount(request.Mobile)
	if err != nil {
		log.Error("account 查询失败 mobile", request.Mobile)
		return reqres.LoginResponse{}, err
	}
	if enPassword != account.Password {
		return reqres.LoginResponse{}, errors.New("账号或密码错误")
	}
	token, tokenErr := utils.GenerateToken("userId", account.UserId, config.TokenSecret)
	if tokenErr != nil {
		log.Error("token获取失败 mobile = " + request.Mobile + " userId = ", account.UserId)
		return reqres.LoginResponse{}, tokenErr
	}
	return reqres.LoginResponse{Token:token}, nil
}

func getAccount(mobile string) (model.Account, error) {
	key := cachekey.Account(mobile)
	begin := time.Now().UnixNano() / 1e6
	value, err := config.RedisClient.Get(key).Result()
	end := time.Now().UnixNano() / 1e6
	log.Info(end - begin)
	if  !utils.StringIsEmpty(value) {
		bytes := []byte(value)
		var cacheAccount model.Account
		// ---------将value转化为account结构体-------------
		err := json.Unmarshal(bytes, &cacheAccount)
		if err != nil {
			log.Error("redis 解析 account 失败 key", key)
			return model.Account{}, err
		}
		log.Info("redis拿到")
		return cacheAccount, nil
	}
	account, err := dao.FindAccountByMobile(mobile)
	if err != nil {
		log.Error("数据库查询失败！FindAccountByMobile mobile = " + mobile)
		return account, err
	}
	log.Info("数据库拿到")
	// -------将account结构体转化为json string-------------
	bytes, err := json.Marshal(account)
	if err != nil {
		log.Error("序列化失败 key + " + key)
		return model.Account{}, err
	}
	strJsonAccount := string(bytes)
	redisErr := config.RedisClient.Set(key, strJsonAccount, config.RedisExpire).Err()
	if redisErr != nil {
		log.Error("redis 查询失败 mobile = " + mobile)
		return model.Account{}, redisErr
	}
	return account, err
}
