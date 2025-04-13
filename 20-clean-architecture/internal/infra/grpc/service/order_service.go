package service

import (
	"context"

	"github.com/rafaelpapastamatiou/goexpert/20-clean-architecture/internal/infra/grpc/pb"
	"github.com/rafaelpapastamatiou/goexpert/20-clean-architecture/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrdersUseCase  usecase.ListOrdersUseCase
}

func NewOrderService(
	createOrderUseCase usecase.CreateOrderUseCase,
	listOrdersUseCase usecase.ListOrdersUseCase,
) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		ListOrdersUseCase:  listOrdersUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.Order, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.Order{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrders(ctx context.Context, in *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	dtos, err := s.ListOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}

	orders := make([]*pb.Order, len(dtos))

	for i, dto := range dtos {
		orders[i] = &pb.Order{
			Id:         dto.ID,
			Price:      float32(dto.Price),
			Tax:        float32(dto.Tax),
			FinalPrice: float32(dto.FinalPrice),
		}
	}

	return &pb.ListOrdersResponse{
		Orders: orders,
	}, nil
}
