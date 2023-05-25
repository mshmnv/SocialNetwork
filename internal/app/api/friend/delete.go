package friend

import (
	"context"

	desc "github.com/mshmnv/SocialNetwork/pkg/api/friend"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Delete реализует /friend/delete/{user_id}
func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*desc.DeleteResponse, error) {

	if req.GetUserId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "Invalid user id")
	}
	if req.GetFriendId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "Invalid friend id")
	}

	if err := i.friendService.DeleteFriend(req.GetFriendId(), req.GetUserId()); err != nil {
		return nil, err
	}

	return &desc.DeleteResponse{}, nil
}
