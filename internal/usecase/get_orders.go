package usecase

import (
	"github.com/Sotnasjeff/clean-arch-api/internal/entity"
)

type ListOrdersOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type GetOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewGetOrdersUseCase(
	orderRepository entity.OrderRepositoryInterface,
) *GetOrdersUseCase {
	return &GetOrdersUseCase{
		OrderRepository: orderRepository,
	}
}

func (g *GetOrdersUseCase) GetOrderByID(id string) (ListOrdersOutputDTO, error) {
	order, err := g.OrderRepository.SelectById(id)
	if err != nil {
		return ListOrdersOutputDTO{}, err
	}
	return ListOrdersOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}

func (g *GetOrdersUseCase) ListAllOrders() ([]ListOrdersOutputDTO, error) {
	orders, err := g.OrderRepository.Select()
	if err != nil {
		return []ListOrdersOutputDTO{}, nil
	}
	var listOrders []ListOrdersOutputDTO
	for _, order := range orders {
		o := ListOrdersOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
		listOrders = append(listOrders, o)
	}

	return listOrders, nil
}
