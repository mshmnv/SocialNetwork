package user

import (
	"context"

	userDS "github.com/mshmnv/SocialNetwork/internal/app/service/user/datastruct"
	desc "github.com/mshmnv/SocialNetwork/pkg/api/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Login реализует /login
func (i *Implementation) Login(ctx context.Context, req *desc.LoginRequest) (*desc.LoginResponse, error) {
	if req.GetId() == 0 || req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "Невалидные данные")
	}

	data := &userDS.LoginData{
		ID:       req.GetId(),
		Password: req.GetPassword(),
	}
	if err := i.userService.Login(data); err != nil {
		return nil, err
	}

	return &desc.LoginResponse{}, nil
}
