package cache

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/mshmnv/SocialNetwork/internal/app/service/post/datastruct"
	"github.com/mshmnv/SocialNetwork/internal/pkg/redis"
	logger "github.com/sirupsen/logrus"
)

const (
	expiration = 72 * time.Hour
)

type IFriendService interface {
	GetUserFriends(userID uint64) ([]uint64, error)
}

type PostCache struct {
	posts     chan *datastruct.Post
	userPosts datastruct.UserPosts
}

func NewPostCacheJob(ctx context.Context, postService IPostService) *PostCache {
	p := &PostCache{
		posts:     make(chan *datastruct.Post),
		userPosts: datastruct.NewUserPosts(),
	}
	go p.StartPostCacheJob(ctx)
	go StartPostUpdateJob(p, postService)
	return p
}

func (p *PostCache) StartPostCacheJob(ctx context.Context) {
	for post := range p.posts {
		if post.IsDeleted {
			if err := p.deletePostFromCache(ctx, post); err != nil {
				logger.Errorf("Error deleting post from cache: %v", err)
			}
			continue
		}
		if err := p.addPostToRedis(ctx, post); err != nil {
			logger.Errorf("Error adding post to cache: %v", err)
		}
	}
}

// UpdatePost обновление постов в кеш в редисе, вызывается в джобе
func (p *PostCache) UpdatePost(post *datastruct.Post) {
	p.posts <- post
}

func createPostRecord(post *datastruct.Post) (key, value string, err error) {
	jsonPost, err := json.Marshal(post)
	if err != nil {
		return "", "", err
	}
	return strconv.Itoa(int(post.PostID)), string(jsonPost), nil
}

func (p *PostCache) addPostToRedis(ctx context.Context, post *datastruct.Post) error {
	key, value, err := createPostRecord(post)
	if err != nil {
		return err
	}

	if err = redis.GetRedis(ctx).Set(key, value, expiration).Err(); err != nil {
		return err
	}
	p.userPosts.AddUserPost(post.AuthorID, post.PostID)
	logger.Infof("POST IS ADDED TO REDIS: key: %s, value: %s", key, value)
	return nil
}

func (p *PostCache) deletePostFromCache(ctx context.Context, post *datastruct.Post) error {
	if err := redis.GetRedis(ctx).Del(strconv.Itoa(int(post.PostID))).Err(); err != nil {
		return err
	}
	p.userPosts.DeleteUserPost(post.AuthorID, post.PostID)
	return nil
}

func (p *PostCache) GetFeedFromCache(ctx context.Context, userIDs []uint64) ([]datastruct.Post, error) {

	postsID := p.userPosts.GetPostsForUsers(userIDs)
	if len(postsID) == 0 {
		logger.Infof("No posts from users %v", userIDs)
		return nil, nil
	}

	result := make([]datastruct.Post, 0, len(postsID))

	for _, postID := range postsID {
		res, err := redis.GetRedis(ctx).Get(strconv.FormatUint(postID, 10)).Result()
		if err != nil {
			logger.Errorf("Error getting posts from redis: %v", err)
			continue
		}
		post := datastruct.Post{}
		err = json.Unmarshal([]byte(res), &post)
		if err != nil {
			logger.Errorf("Error unmarshalling posts from redis: %v", err)
			continue
		}
		result = append(result, post)
	}
	return result, nil
}
