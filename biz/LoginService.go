package biz

import (
	"github.com/impte/landisland/constant"
	"github.com/impte/landisland/model"
	"github.com/impte/landisland/reqres"
)

func Login(loginRequest reqres.LoginRequest) reqres.LoginResponse {
	account := model.Account{
		UserId: "123456789",
		Username: "徐通",
		Password: "789456",
		Mobile: "17725398992",
		Status: 1,
	}
	err := account.Create(constant.Database)
	if err != nil {
		panic(err.Error())
	}
	return reqres.LoginResponse{}
}
