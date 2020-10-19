package service

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/iamMarkchu/iron/app/http/request"
	"github.com/iamMarkchu/iron/app/model"
	"github.com/iamMarkchu/iron/core/auth/jwt"
	"github.com/iamMarkchu/iron/core/store/orm"
)

type userService struct {

}

func (s *userService) Register(ctx context.Context, req request.RegisterReq) (int64, error)  {
	var (
		o = orm.GetInstance()
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

func (s *userService) Login(ctx *gin.Context, req request.LoginReq) (auth jwt.Auth, err error) {
	var (
		o = orm.GetInstance()
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

func NewUserService() *userService {
	return &userService{}
}
