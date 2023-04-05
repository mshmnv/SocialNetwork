package user

import (
	"context"

	desc "github.com/mshmnv/SocialNetwork/pkg/api/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetUser реализует /user/get/{id}
func (i *Implementation) GetUser(ctx context.Context, req *desc.GetUserRequest) (*desc.GetUserResponse, error) {

	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "Невалидные данные")
	}

	user, err := i.userService.GetUser(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &desc.GetUserResponse{
		FirstName:  user.FirstName,
		SecondName: user.SecondName,
		Age:        user.Age,
		Birthdate:  user.BirthDate,
		Biography:  user.Biography,
		City:       user.City,
	}, nil
}
