package service

import (
	"context"
	"google.golang.org/grpc"
	"zhuhui.com/microkit/grpc_pb/pb"
)

type ProdService struct {
}

func (this *ProdService) GetProdStock(ctx context.Context, in *proc_pb_pb.ProdRequest, opts ...grpc.CallOption) (*proc_pb_pb.ProdResponse, error) {

	return &proc_pb_pb.ProdResponse{ProdStock: 20}, nil
}
