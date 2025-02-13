package controller

import (
	"context"

	productdomain "github.com/haohmaru3000/explicit_architecture/module/product/domain"
)

type CreateProductUseCase interface {
	CreateProduct(ctx context.Context, prod *productdomain.ProductCreationDTO) error
}

type APIController struct {
	createUseCase CreateProductUseCase
}

func NewAPIController(createUseCase CreateProductUseCase) APIController {
	return APIController{createUseCase: createUseCase}
}
