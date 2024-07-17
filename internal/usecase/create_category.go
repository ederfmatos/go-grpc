package usecase

import (
	"context"
	"go-grpc/internal/entity"
	"go-grpc/internal/repository"
)

type Input struct {
	Context     context.Context
	Name        string
	Description string
}

type Output struct {
	Id string `json:"id"`
}

type UseCase struct {
	categoryRepository repository.CategoryRepository
}

func New(categoryRepository repository.CategoryRepository) *UseCase {
	return &UseCase{categoryRepository: categoryRepository}
}

func (useCase *UseCase) Execute(input Input) (*Output, error) {
	category := entity.NewCategory(input.Name, input.Description)
	err := useCase.categoryRepository.Create(input.Context, category)
	if err != nil {
		return nil, err
	}
	return &Output{Id: category.ID}, nil
}
