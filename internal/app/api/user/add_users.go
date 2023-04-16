package user

import (
	"context"

	desc "github.com/mshmnv/SocialNetwork/pkg/api/user"
)

// AddUsers реализует /add-users
func (i *Implementation) AddUsers(ctx context.Context, req *desc.AddUsersRequest) (*desc.AddUsersResponse, error) {

	i.userService.AddUsers()

	return &desc.AddUsersResponse{}, nil
}
