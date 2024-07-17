package rpc

import (
	"context"
	"go-grpc/internal/pb"
	"go-grpc/internal/repository"
	"go-grpc/internal/usecase"
	"io"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	createCategoryUseCase *usecase.UseCase
	categoryRepository    repository.CategoryRepository
}

func NewCategoryService(categoryRepository repository.CategoryRepository) pb.CategoryServiceServer {
	return &CategoryService{
		createCategoryUseCase: usecase.New(categoryRepository),
		categoryRepository:    categoryRepository,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, request *pb.CreateCategoryRequest) (*pb.CreateCategoryResponse, error) {
	input := usecase.Input{
		Context:     ctx,
		Name:        request.Name,
		Description: request.Description,
	}
	output, err := c.createCategoryUseCase.Execute(input)
	if err != nil {
		return nil, err
	}
	return &pb.CreateCategoryResponse{Id: output.Id}, nil
}

func (c *CategoryService) ListCategories(ctx context.Context, _ *pb.Empty) (*pb.ListCategoriesResponse, error) {
	output, err := c.categoryRepository.List(ctx)
	if err != nil {
		return nil, err
	}
	categories := make([]*pb.Category, len(output))
	for i, category := range output {
		categories[i] = &pb.Category{
			Id:          category.Id,
			Name:        category.Name,
			Description: *category.Description,
		}
	}
	return &pb.ListCategoriesResponse{Categories: categories}, nil
}

func (c *CategoryService) GetCategoryById(ctx context.Context, request *pb.GetCategoryByIdRequest) (*pb.GetCategoryByIdResponse, error) {
	output, err := c.categoryRepository.FindById(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetCategoryByIdResponse{
		Id:          output.ID,
		Name:        output.Name,
		Description: output.Description,
	}, nil
}

func (c *CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) error {
	ctx := stream.Context()

	response := &pb.ListCategoriesResponse{
		Categories: make([]*pb.Category, 0),
	}

	for {
		request, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(response)
		}
		if err != nil {
			return err
		}

		input := usecase.Input{
			Context:     ctx,
			Name:        request.Name,
			Description: request.Description,
		}
		output, err := c.createCategoryUseCase.Execute(input)
		if err != nil {
			return err
		}
		response.Categories = append(response.Categories, &pb.Category{
			Id:          output.Id,
			Name:        input.Name,
			Description: input.Description,
		})
	}
}

func (c *CategoryService) CreateCategoryStreamBidirectional(stream pb.CategoryService_CreateCategoryStreamBidirectionalServer) error {
	ctx := stream.Context()

	for {
		request, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		input := usecase.Input{
			Context:     ctx,
			Name:        request.Name,
			Description: request.Description,
		}
		output, err := c.createCategoryUseCase.Execute(input)
		if err != nil {
			return err
		}
		category := &pb.Category{
			Id:          output.Id,
			Name:        input.Name,
			Description: input.Description,
		}
		err = stream.Send(category)
		if err != nil {
			return nil
		}
	}
}
