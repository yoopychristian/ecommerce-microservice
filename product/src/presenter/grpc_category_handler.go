package presenter

import (
	"errors"

	pb "ecommerce-microservice/product/protogo/category"
	"ecommerce-microservice/product/src/model"
	"ecommerce-microservice/product/src/usecase"

	"golang.org/x/net/context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//GrpcCategoryHandler
type GrpcCategoryHandler struct {
	categoryUseCase usecase.CategoryUseCase
}

// NewGrpcCategoryHandler
func NewGrpcCategoryHandler(categoryUseCase usecase.CategoryUseCase) *GrpcCategoryHandler {
	return &GrpcCategoryHandler{categoryUseCase}
}

//FindByID
func (h *GrpcCategoryHandler) FindByID(ctx context.Context, arg *pb.CategoryQueryRequest) (*pb.CategoryResponse, error) {

	id := arg.ID

	categoryResult := <-h.categoryUseCase.FindByID(int(id))

	if categoryResult.Error != nil {
		return nil, status.Error(codes.InvalidArgument, categoryResult.Error.Error())
	}

	category, ok := categoryResult.Result.(*model.Category)

	if !ok {
		err := errors.New("Result is not Category")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	categoryResponse := &pb.CategoryResponse{
		ID:          int32(category.ID),
		Name:        category.Name,
		Description: category.Description,
		Image:       category.Image,
	}

	return categoryResponse, nil
}

// FindAll
func (h *GrpcCategoryHandler) FindAll(arg *pb.CategoryQueryRequest, stream pb.CategoryService_FindAllServer) error {

	categoryResult := <-h.categoryUseCase.FindAll()

	if categoryResult.Error != nil {
		return status.Error(codes.InvalidArgument, categoryResult.Error.Error())
	}

	categories, ok := categoryResult.Result.(model.Categories)

	if !ok {
		err := errors.New("Result is not Categories")
		return status.Error(codes.InvalidArgument, err.Error())
	}

	for _, category := range categories {
		categoryResponse := &pb.CategoryResponse{
			ID:          int32(category.ID),
			Name:        category.Name,
			Description: category.Description,
			Image:       category.Image,
		}

		if err := stream.Send(categoryResponse); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}

	return nil
}
