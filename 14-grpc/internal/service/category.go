package service

import (
	"context"
	"io"

	"github.com/rafaelpapastamatiou/goexpert/14-grpc/internal/database"
	"github.com/rafaelpapastamatiou/goexpert/14-grpc/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB *database.Category
}

func NewCategoryService(db *database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: db,
	}
}

func (s *CategoryService) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := s.CategoryDB.Create(req.Name, req.Description)
	if err != nil {
		return nil, err
	}

	res := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return res, nil
}

func (s *CategoryService) ListCategories(ctx context.Context, req *pb.Blank) (*pb.CategoryList, error) {
	categories, err := s.CategoryDB.List()
	if err != nil {
		return nil, err
	}

	results := make([]*pb.Category, len(categories))

	for i, category := range categories {
		results[i] = &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		}
	}

	res := &pb.CategoryList{
		Categories: results,
	}

	return res, nil
}

func (s *CategoryService) GetCategory(ctx context.Context, req *pb.GetCategoryRequest) (*pb.Category, error) {
	category, err := s.CategoryDB.FindById(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}

func (s *CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) error {
	categories := &pb.CategoryList{}

	for {
		createCategoryReq, err := stream.Recv()

		// Check if the stream has ended
		if err == io.EOF {
			// Send all created categories at once
			return stream.SendAndClose(categories)
		}

		if err != nil {
			return err
		}

		category, err := s.CategoryDB.Create(createCategoryReq.Name, createCategoryReq.Description)
		if err != nil {
			return err
		}

		categories.Categories = append(categories.Categories, &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}
}

func (s *CategoryService) CreateCategoryBidirectionalStream(stream pb.CategoryService_CreateCategoryBidirectionalStreamServer) error {
	for {
		createCategoryReq, err := stream.Recv()

		// Check if the stream has ended
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		category, err := s.CategoryDB.Create(createCategoryReq.Name, createCategoryReq.Description)
		if err != nil {
			return err
		}

		res := &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		}

		// Send the created category back to the client
		// This is a bidirectional stream, so we can send messages back to the client
		// while still receiving messages from the client
		// in the same stream
		err = stream.Send(res)
		if err != nil {
			return err
		}
	}
}
