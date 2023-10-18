package post

import (
	"github.com/mshmnv/SocialNetwork/internal/app/service/post/datastruct"
	postService "github.com/mshmnv/SocialNetwork/internal/app/service/post/service"
	"github.com/mshmnv/SocialNetwork/internal/pkg/postgres"
	"github.com/mshmnv/SocialNetwork/internal/pkg/redis"
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

func NewPostAPI(db *postgres.DB, redisClient *redis.Client) *Implementation {

	return &Implementation{
		postService: postService.BuildService(db, redisClient),
	}
}
