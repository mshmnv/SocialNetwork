package post

import (
	"context"

	"github.com/mshmnv/SocialNetwork/internal/app/service/post/datastruct"
	postService "github.com/mshmnv/SocialNetwork/internal/app/service/post/service"
	postAPI "github.com/mshmnv/SocialNetwork/pkg/api/post"
)

type IPostService interface {
	Create(userID uint64, text string) error
	Update(userID, postID uint64, text string) error
	Delete(userID, postID uint64) error
	Feed(userID, offset, limit uint64) ([]datastruct.Post, error)
	Get(postID uint64) (*datastruct.Post, error)
	AddPosts(userID uint64) error
}

type Implementation struct {
	postService IPostService
	postAPI.UnimplementedPostAPIServer
}

func NewPostAPI(ctx context.Context) *Implementation {

	return &Implementation{
		postService: postService.BuildService(ctx),
	}
}
