package service

import (
	"context"

	"github.com/Sotnasjeff/clean-arch-api/internal/infra/grpc/pb"
	"github.com/Sotnasjeff/clean-arch-api/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	GetOrdersUseCase   usecase.GetOrdersUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, getOrdersUseCase usecase.GetOrdersUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		GetOrdersUseCase:   getOrdersUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListAllOrders(ctx context.Context, in *pb.Blank) (*pb.ListOrdersResponse, error) {
	ordersOutput, err := s.GetOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}
	var listOrders []*pb.GetOrderResponse
	for _, order := range ordersOutput {
		o := &pb.GetOrderResponse{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		}
		listOrders = append(listOrders, o)
	}
	return &pb.ListOrdersResponse{Orders: listOrders}, nil
}
func (s *OrderService) GetOrderById(ctx context.Context, in *pb.GetOrderByIdRequest) (*pb.GetOrderResponse, error) {
	order, err := s.GetOrdersUseCase.GetOrderByID(in.Id)
	if err != nil {
		return nil, err
	}

	orderResponse := &pb.GetOrderResponse{
		Id:         order.ID,
		Price:      float32(order.Price),
		Tax:        float32(order.Tax),
		FinalPrice: float32(order.FinalPrice),
	}

	return orderResponse, nil
}
