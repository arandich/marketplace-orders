package service

import (
	"context"
	pb "github.com/arandich/marketplace-proto/api/proto/services"
)

type OrdersRepository interface {
	GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error)
	InitOrder(ctx context.Context, req *pb.InitOrderRequest) (*pb.InitOrderResponse, error)
}

var _ OrdersRepository = (*OrdersService)(nil)

type OrdersService struct {
	pb.UnimplementedOrderServiceServer
	repository OrdersRepository
}

func NewIdService(repository OrdersRepository) OrdersService {
	return OrdersService{
		repository: repository,
	}
}

func (s OrdersService) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	return s.repository.GetOrder(ctx, req)
}

func (s OrdersService) InitOrder(ctx context.Context, req *pb.InitOrderRequest) (*pb.InitOrderResponse, error) {
	return s.repository.InitOrder(ctx, req)
}
