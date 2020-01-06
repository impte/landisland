package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//go:generate goqueryset -in Account.go
// gen:qs
type Account struct {
	gorm.Model

	BaseEntity
	UserId string
	Username string
	Password string
	Mobile string
	Status int
}
