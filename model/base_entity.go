package model

type BaseEntity struct {
	Id         int32 `json:"id"`
	CreateTime int32 `json:"createTime"`
	UpdateTime int32 `json:"updateTime"`
}
