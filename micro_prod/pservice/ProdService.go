package pservice

import (
	"context"
	"strconv"
	"time"
	. "zhuhui.com/microkit/micro_pb/pb"
)

type ProdServiceImpl struct {
}

func (ps *ProdServiceImpl) GetProdDetail(ctx context.Context, req *ProdsRequest, rsp *ProdDetailResponse) error {
	rsp.Data = newProd(req.ProdId, "test prod")
	return nil
}

func newProd(id int32, pname string) *ProdModel {
	return &ProdModel{ProdID: id, ProdName: pname}
}

func (ps *ProdServiceImpl) GetProdsList(ctx context.Context, req *ProdsRequest, resp *ProdListResponse) error {
	time.Sleep(time.Second * 3)
	models := make([]*ProdModel, 0)
	var i int32
	for i = 0; i < req.Size; i++ {
		models = append(models, newProd(100+i, "prod-"+strconv.Itoa(100+int(i))))
	}
	resp.Data = models
	return nil
}
