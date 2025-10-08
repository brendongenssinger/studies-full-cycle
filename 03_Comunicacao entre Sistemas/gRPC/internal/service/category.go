package service

import (
	"context"

	"github.com/brendonmascarenhas/gqlgen-todos/internal/database"
	"github.com/brendonmascarenhas/gqlgen-todos/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	categoryDB database.Category
}

func NewCategoryService(db database.Category) *CategoryService {
	return &CategoryService{categoryDB: db}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.categoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}
	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
	return &pb.CategoryResponse{Category: categoryResponse}, nil
}
