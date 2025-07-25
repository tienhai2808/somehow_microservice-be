package service

import (
	"context"

	"github.com/SomeHowMicroservice/shm-be/services/product/model"
	"github.com/SomeHowMicroservice/shm-be/services/product/protobuf"
)

type ProductService interface {
	CreateCategory(ctx context.Context, req *protobuf.CreateCategoryRequest) (*model.Category, error)

	GetCategoryTree(ctx context.Context) ([]*model.Category, error) 

	CreateProduct(ctx context.Context, req *protobuf.CreateProductRequest) (*model.Product, error)

	GetProductBySlug(ctx context.Context, slug string) (*model.Product, error)
}