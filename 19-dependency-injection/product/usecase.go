package product

type ProductUsecase struct {
	productRepository ProductRepositoryInterface
}

func NewProductUsecase(productRepository ProductRepositoryInterface) *ProductUsecase {
	return &ProductUsecase{productRepository: productRepository}
}

func (u *ProductUsecase) GetProductByID(id int) (*Product, error) {
	product, err := u.productRepository.GetProductByID(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}
