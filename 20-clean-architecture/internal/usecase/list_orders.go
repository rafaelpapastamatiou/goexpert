package usecase

import (
	"github.com/rafaelpapastamatiou/goexpert/20-clean-architecture/internal/entity"
)

type ListOrdersOrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(orderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: orderRepository,
	}
}

func (u *ListOrdersUseCase) Execute() ([]ListOrdersOrderOutputDTO, error) {
	orders, err := u.OrderRepository.FindAll()
	if err != nil {
		return nil, err
	}

	dtos := make([]ListOrdersOrderOutputDTO, len(orders))

	for i, order := range orders {
		dtos[i] = ListOrdersOrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.Price + order.Tax,
		}
	}

	return dtos, nil
}
