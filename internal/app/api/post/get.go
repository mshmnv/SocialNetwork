package post

import (
	"context"

	desc "github.com/mshmnv/SocialNetwork/pkg/api/post"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Get реализует /post/get/{id}
func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {

	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "Invalid post id")
	}

	post, err := i.postService.Get(req.GetId())
	if err != nil {
		return nil, err
	}

	return &desc.GetResponse{Post: &desc.Post{
		Id:       post.PostID,
		Text:     post.Text,
		AuthorId: post.AuthorID,
	}}, nil
}
