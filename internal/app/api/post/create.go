package post

import (
	"context"

	desc "github.com/mshmnv/SocialNetwork/pkg/api/post"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create реализует /post/create
func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {

	userID := ctx.Value("user_id").(uint64)
	if userID == 0 {
		return nil, status.Error(codes.InvalidArgument, "Invalid user id")
	}
	if req.GetText() == "" {
		return nil, status.Error(codes.InvalidArgument, "Empty post")
	}

	if err := i.postService.Create(userID, req.GetText()); err != nil {
		return nil, err
	}

	return &desc.CreateResponse{}, nil
}
