package service

import (
	"context"
	"strconv"
	. "zhuhui.com/microkit/micro/pb"
)

type TestServiceImpl struct {
}

func (this *TestServiceImpl) Call(ctx context.Context, req *TestRequest, rep *TestResponse) error {
	rep.Data = "test" + strconv.Itoa(int(req.Id))
	return nil
}
