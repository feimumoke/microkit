package service

import (
	"context"
	"time"
	"zhuhui.com/microkit/micro/appinit"
	"zhuhui.com/microkit/micro/dbmodel"
	. "zhuhui.com/microkit/micro/pb"
)

type UserServiceImpl struct {
}

/**
apigw.sh

http://localhost:8080/test/UserService/UserReg
*/

func (this *UserServiceImpl) UserReg(ctx context.Context, user *UserModel, res *RegResponse) error {
	users := dbmodel.Users{UserName: user.Username, UserPwd: user.UserPwd, UserDate: time.Now()}
	if err := appinit.GetDB().Create(&users).Error; err != nil {
		res.Message = err.Error()
		res.Status = "error"
	} else {
		res.Message = ""
		res.Status = "success"
	}
	return nil
}
