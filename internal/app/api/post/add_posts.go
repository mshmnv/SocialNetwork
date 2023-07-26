package post

import (
	"context"

	desc "github.com/mshmnv/SocialNetwork/pkg/api/post"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AddPosts реализует /add-posts
func (i *Implementation) AddPosts(ctx context.Context, req *desc.AddPostsRequest) (*desc.AddPostsResponse, error) {

	userId := ctx.Value("user_id").(uint64)
	if userId == 0 {
		return nil, status.Error(codes.InvalidArgument, "Invalid user id")
	}

	if err := i.postService.AddPosts(userId); err != nil {
		return nil, err
	}
	return &desc.AddPostsResponse{}, nil
}
