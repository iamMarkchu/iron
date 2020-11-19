package core

import (
	"github.com/iamMarkchu/iron/app/service"
	"github.com/iamMarkchu/iron/core/auth/jwt"
	"github.com/iamMarkchu/iron/core/config"
	"github.com/iamMarkchu/iron/core/server"
	"github.com/iamMarkchu/iron/core/store/orm"
)

func Init() {
	config.Init()
	server.Init()
	service.Init()
	orm.Init()
	jwt.Init()
}
