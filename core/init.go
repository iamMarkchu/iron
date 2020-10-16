package core

import (
	"github.com/iamMarkchu/iron/core/config"
	"github.com/iamMarkchu/iron/core/server"
	"github.com/iamMarkchu/iron/core/store/orm"
)

func Init()  {
	server.Init()
	config.Init()
	orm.Init()
}
