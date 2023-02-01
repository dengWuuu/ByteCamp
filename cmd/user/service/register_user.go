package service

import (
	"context"
	"douyin/dal/db"
	"douyin/kitex_gen/user"
	"douyin/pkg/bcrypt"
	"douyin/pkg/errno"

	"github.com/cloudwego/kitex/pkg/klog"
)

type RegisterService struct {
	ctx context.Context
}

func NewRegisterService(ctx context.Context) RegisterService {
	return RegisterService{ctx: ctx}
}
func (registerService RegisterService) Register(req *user.DouyinUserRegisterRequest) error {
	userLists, err := db.GetUsersByUserName(req.GetUsername())
	//首先查询数据库中是否有该用户
	if err != nil {
		klog.Fatalf("数据库中根据用户查找用户名报错")
		return err
	}
	if len(userLists) != 0 {
		return errno.ErrUserAlreadyExist
	}
	//加密密码信息
	p, err := bcrypt.PasswordHash(req.Password)
	if err != nil {
		klog.Fatalf("加密密码出现异常")
		return err
	}
	//不存在该用户 直接插入该用户数据
	return db.CreateUser(&db.User{Name: req.Username, Password: p})
}
