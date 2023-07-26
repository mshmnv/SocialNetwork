package friend

import (
	"context"

	desc "github.com/mshmnv/SocialNetwork/pkg/api/friend"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Delete реализует /friend/delete/{user_id}
func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*desc.DeleteResponse, error) {

	userID := ctx.Value("user_id").(uint64)
	if userID == 0 {
		return nil, status.Error(codes.InvalidArgument, "Invalid user id")
	}

	if req.GetUserId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "Invalid friend id")
	}

	if err := i.friendService.DeleteFriend(req.GetUserId(), userID); err != nil {
		return nil, err
	}

	return &desc.DeleteResponse{}, nil
}
