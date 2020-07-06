package service

import (
	"context"
	. "zhuhui.com/microkit/micro/pb"
)

type UserServiceImpl struct {
}

func (this *UserServiceImpl) UserReg(ctx context.Context, req *UserModel, res *RegResponse) error {
	res.Message = ""
	res.Status = "success"
	return nil
}
