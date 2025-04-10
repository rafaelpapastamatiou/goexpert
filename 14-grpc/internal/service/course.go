package service

import (
	"context"

	"github.com/rafaelpapastamatiou/goexpert/14-grpc/internal/database"
	"github.com/rafaelpapastamatiou/goexpert/14-grpc/internal/pb"
)

type CourseService struct {
	pb.UnimplementedCourseServiceServer
	CourseDB *database.Course
}

func NewCourseService(db *database.Course) *CourseService {
	return &CourseService{
		CourseDB: db,
	}
}

func (s *CourseService) CreateCourse(ctx context.Context, req *pb.CreateCourseRequest) (*pb.Course, error) {
	course, err := s.CourseDB.Create(req.Name, req.Description, req.CategoryId)
	if err != nil {
		return nil, err
	}

	res := &pb.Course{
		Id:          course.ID,
		Name:        course.Name,
		Description: course.Description,
		CategoryId:  course.CategoryID,
	}

	return res, nil
}
