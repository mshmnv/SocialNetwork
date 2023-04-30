package user

import (
	"context"

	"github.com/mshmnv/SocialNetwork/internal/app/service/user/datastruct"
	userService "github.com/mshmnv/SocialNetwork/internal/app/service/user/service"
	userAPI "github.com/mshmnv/SocialNetwork/pkg/api/user"
)

type IUserService interface {
	Register(data *datastruct.User) error
	GetUser(id uint64) (*datastruct.User, error)
	Login(data *datastruct.LoginData) error
	AddUsers()
	Search(firstName, secondName string) ([]datastruct.User, error)
}

type Implementation struct {
	userService IUserService
	userAPI.UnimplementedUserAPIServer
}

func NewUserAPI(ctx context.Context) *Implementation {
	return &Implementation{
		userService: userService.BuildService(ctx),
	}
}
