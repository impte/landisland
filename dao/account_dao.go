package dao

import (
	"github.com/impte/landisland/config"
	"github.com/impte/landisland/model"
)

func FindAccountByMobile(mobile string) (model.Account, error) {
	account := model.Account{}
	err := model.NewAccountQuerySet(config.Database).
		MobileEq(mobile).
		StatusEq(1).
		Select(
			model.AccountDBSchema.UserId,
			model.AccountDBSchema.Username,
			model.AccountDBSchema.Password,
			model.AccountDBSchema.Mobile,
		).
		One(&account)
	return account, err
}
