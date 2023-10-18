package service

import (
	"bufio"
	"os"
	"time"

	friendService "github.com/mshmnv/SocialNetwork/internal/app/service/friend/service"
	"github.com/mshmnv/SocialNetwork/internal/app/service/post/cache"
	"github.com/mshmnv/SocialNetwork/internal/app/service/post/datastruct"
	"github.com/mshmnv/SocialNetwork/internal/app/service/post/repository"
	"github.com/mshmnv/SocialNetwork/internal/pkg/postgres"
	"github.com/mshmnv/SocialNetwork/internal/pkg/redis"
	logger "github.com/sirupsen/logrus"
)

type IRepository interface {
	Create(userID uint64, text string) error
	Update(userID, postID uint64, text string) error
	Delete(userID, postID uint64) error
	Get(postID uint64) (*datastruct.Post, error)
	GetPostsAfterDate(date time.Time) ([]datastruct.Post, error)
}

type IFriendService interface {
	GetUserFriends(userID uint64) ([]uint64, error)
}

type IPostCache interface {
	GetFeedFromCache(userID []uint64) ([]datastruct.Post, error)
}

type Service struct {
	repository    IRepository
	postCache     IPostCache
	friendService IFriendService
}

func BuildService(db *postgres.DB, redisClient *redis.Client) *Service {
	s := &Service{
		repository:    repository.NewRepository(db),
		friendService: friendService.BuildService(db),
	}

	s.postCache = cache.NewPostCacheJob(s, redisClient)
	return s
}

func (s *Service) Create(userID uint64, text string) error {
	return s.repository.Create(userID, text)
}

func (s *Service) Update(userID, postID uint64, text string) error {
	return s.repository.Update(userID, postID, text)
}

func (s *Service) Delete(userID, postID uint64) error {
	return s.repository.Delete(userID, postID)
}

func (s *Service) Feed(userID, offset, limit uint64) ([]datastruct.Post, error) {

	friendsIDs, err := s.friendService.GetUserFriends(userID)
	if err != nil {
		logger.Errorf("Error getting user's friends: %v", err)
		return nil, err
	}
	if len(friendsIDs) == 0 {
		logger.Infof("No friend for user %v", userID)
		return nil, nil
	}

	posts, err := s.postCache.GetFeedFromCache(friendsIDs)
	if err != nil {
		logger.Errorf("Error getting posts from cache: %v", err)
		return nil, err
	}

	if len(posts) < int(offset) {
		return []datastruct.Post{}, nil
	}

	if len(posts) < int(offset+limit) {
		return posts[offset:], nil
	}

	return posts[offset : offset+limit], nil
}

func (s *Service) Get(postID uint64) (*datastruct.Post, error) {
	return s.repository.Get(postID)
}

func (s *Service) GetPostsAfterDate(date time.Time) ([]datastruct.Post, error) {
	return s.repository.GetPostsAfterDate(date)
}

func (s *Service) AddPosts(userID uint64) error {
	f, err := os.Open("info/testing/posts.txt")
	if err != nil {
		logger.Errorf("Error reading file: %v", err)
	}

	scanner := bufio.NewScanner(f)
	go func() {
		defer f.Close()
		for scanner.Scan() {
			if err := s.repository.Create(userID, scanner.Text()); err != nil {
				logger.Errorf("Error creating posts: %v", err)
			}
		}
		if err := scanner.Err(); err != nil {
			logger.Errorf("Error creating posts: %v", err)
		}
		logger.Info("Posts are successfully added")
	}()
	return nil
}
