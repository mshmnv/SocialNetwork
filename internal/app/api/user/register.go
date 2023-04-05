package user

import (
	"context"

	userDS "github.com/mshmnv/SocialNetwork/internal/app/service/user/datastruct"
	desc "github.com/mshmnv/SocialNetwork/pkg/api/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Register реализует /user/register
func (i *Implementation) Register(ctx context.Context, req *desc.RegisterRequest) (*desc.RegisterResponse, error) {

	if !isValidData(req) {
		return nil, status.Error(codes.InvalidArgument, "Невалидные данные")
	}

	userData := &userDS.User{
		FirstName:  req.GetFirstName(),
		SecondName: req.GetSecondName(),
		Age:        req.GetAge(),
		BirthDate:  req.GetBirthdate(),
		Biography:  req.GetBiography(),
		City:       req.GetCity(),
		Password:   req.GetPassword(),
	}

	err := i.userService.Register(ctx, userData)
	if err != nil {
		return nil, err
	}

	return &desc.RegisterResponse{}, nil
}

func isValidData(req *desc.RegisterRequest) bool {
	if req.GetFirstName() == "" || req.GetSecondName() == "" ||
		req.GetAge() <= 0 || req.GetBirthdate() == "" ||
		req.GetCity() == "" || req.GetPassword() == "" {
		return false
	}
	return true
}
