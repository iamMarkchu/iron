package service

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/iamMarkchu/iron/app/http/request"
	"github.com/iamMarkchu/iron/app/model"
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

func (s *userService) Login(ctx *gin.Context, req request.LoginReq) error {
	var (
		o = orm.GetInstance()
		userModel model.User
	)
	userModel.UserName = req.UserName
	userModel.Password = req.Password
	// 查询用户
	o.Where("user_name = ?", req.UserName).Find(&userModel)
	if userModel.Id == 0 {
		return errors.New("用户名不存在")
	}
	//校验密码
	if userModel.Password != req.Password {
		return errors.New("密码错误")
	}
	// 返回 token 和 用户id
}

func NewUserService() *userService {
	return &userService{}
}
