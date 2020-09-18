package models

import "time"

type User struct {
	Id          int64
	UserName    string
	NickName    string
	Mobile      string
	Avatar      string
	Description string
	Age         int8
	WxId        string
	Status      int8
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (m *User) TableName() string {
	return "tt_users"
}
