package friend

import (
	"context"

	desc "github.com/mshmnv/SocialNetwork/pkg/api/friend"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Set реализует /friend/set/{user_id}
func (i *Implementation) Set(ctx context.Context, req *desc.SetRequest) (*desc.SetResponse, error) {

	if req.GetUserId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "Invalid user id")

	}
	if req.GetFriendId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "Invalid friend id")
	}

	if req.GetUserId() == req.GetFriendId() {
		return nil, status.Error(codes.InvalidArgument, "User id can not be equal to friend id")
	}

	if err := i.friendService.SendFriendRequest(req.GetFriendId(), req.GetUserId()); err != nil {
		return nil, err
	}
	return &desc.SetResponse{}, nil
}
