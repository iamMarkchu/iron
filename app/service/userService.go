package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/iamMarkchu/iron/app/model"
	"github.com/iamMarkchu/iron/app/model/request"
	"github.com/iamMarkchu/iron/core/auth/jwt"
	"github.com/iamMarkchu/iron/core/store/orm"
	"github.com/silenceper/wechat/v2"

	cache2 "github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram/config"
)

type userService struct {
}

func (s *userService) Register(ctx context.Context, req request.RegisterReq) (int64, error) {
	var (
		o         = orm.GetInstance()
		userModel model.User
	)
	if req.Password != req.RePassword {
		return 0, errors.New("两次密码不一致")
	}
	userModel.UserName = req.UserName
	userModel.Password = req.Password
	o.Create(&userModel)
	return userModel.Id, nil
}

func (s *userService) Login(ctx context.Context, req request.LoginReq) (auth jwt.Auth, err error) {
	var (
		o         = orm.GetInstance()
		userModel model.User
	)
	userModel.UserName = req.UserName
	userModel.Password = req.Password
	// 查询用户
	o.Where("user_name = ?", req.UserName).Find(&userModel)
	if userModel.Id == 0 {
		err = errors.New("用户不存在")
		return
	}
	//校验密码
	if userModel.Password != req.Password {
		err = errors.New("密码错误")
		return
	}

	// 返回 token 和 用户id
	auth, err = jwt.GetToken(int(userModel.Id))
	return
}

func (s *userService) WxLogin(ctx context.Context, code string) {
	wc := wechat.NewWechat()
	cache := cache2.NewRedis(&cache2.RedisOpts{
		Host: "localhost:6379",
	})
	cfg := &config.Config{
		AppID:     "wx4c07089adcd15e4e",
		AppSecret: "a1d281d829928caafaa33ea3f8a8aa12",
		Cache:     cache,
	}
	mini := wc.GetMiniProgram(cfg)
	a := mini.GetAuth()
	res, err := a.Code2Session(code)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("res:", res)
}

func NewUserService() *userService {
	return &userService{}
}
