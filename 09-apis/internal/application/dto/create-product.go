package dto

type CreateProductRequestBodyDTO struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CreateProductResponseDTO struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
