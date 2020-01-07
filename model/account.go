package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//go:generate goqueryset -in account.go
// gen:qs
type Account struct {
	BaseEntity
	UserId   string `json:"userId"`
	Username string `json:"username"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
	Status   int    `json:"status"`
}
