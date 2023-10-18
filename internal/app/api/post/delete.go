package post

import (
	"context"

	desc "github.com/mshmnv/SocialNetwork/pkg/api/post"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Delete реализует /post/delete/{id}
func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*desc.DeleteResponse, error) {

	userID := ctx.Value("user_id").(uint64)
	if userID == 0 {
		return nil, status.Error(codes.InvalidArgument, "Invalid user id")
	}

	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "Invalid post id")
	}

	if err := i.postService.Delete(userID, req.GetId()); err != nil {
		return nil, err
	}

	return &desc.DeleteResponse{}, nil
}
