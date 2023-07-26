package post

import (
	"context"

	desc "github.com/mshmnv/SocialNetwork/pkg/api/post"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Update реализует /post/update/{id}
func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*desc.UpdateResponse, error) {

	userID := ctx.Value("user_id").(uint64)
	if userID == 0 {
		return nil, status.Error(codes.InvalidArgument, "Invalid user id")
	}

	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "Empty post id")
	}

	if req.GetText() == "" {
		return nil, status.Error(codes.InvalidArgument, "Empty post")
	}

	if err := i.postService.Update(userID, req.GetId(), req.GetText()); err != nil {
		return nil, err
	}

	return &desc.UpdateResponse{}, nil
}
