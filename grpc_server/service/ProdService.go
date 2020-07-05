package service

import (
	"context"
	"zhuhui.com/microkit/grpc_pb/pb"
)

type ProdService struct {
}

func (ProdService) GetProdStockList(context.Context, *proc_pb_pb.QuerySize) (*proc_pb_pb.ProdResponseList, error) {
	prods := []*proc_pb_pb.ProdResponse{
		&proc_pb_pb.ProdResponse{ProdStock: 10},

		&proc_pb_pb.ProdResponse{ProdStock: 20},
		&proc_pb_pb.ProdResponse{ProdStock: 30},
		&proc_pb_pb.ProdResponse{ProdStock: 40},
	}
	return &proc_pb_pb.ProdResponseList{Prodres: prods}, nil
}

func (ProdService) GetProdStock(ctx context.Context, in *proc_pb_pb.ProdRequest) (*proc_pb_pb.ProdResponse, error) {
	return &proc_pb_pb.ProdResponse{ProdStock: 20}, nil
}
