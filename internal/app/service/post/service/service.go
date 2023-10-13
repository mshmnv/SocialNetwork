package service

import (
	"bufio"
	"os"
	"time"

	friendService "github.com/mshmnv/SocialNetwork/internal/app/service/friend/service"
	"github.com/mshmnv/SocialNetwork/internal/app/service/post/cache"
	"github.com/mshmnv/SocialNetwork/internal/app/service/post/datastruct"
	"github.com/mshmnv/SocialNetwork/internal/app/service/post/repository"
	"github.com/mshmnv/SocialNetwork/internal/pkg/clients"
	"github.com/mshmnv/SocialNetwork/internal/pkg/postgres"
	"github.com/mshmnv/SocialNetwork/internal/pkg/rabbitmq"
	"github.com/mshmnv/SocialNetwork/internal/pkg/redis"
	logger "github.com/sirupsen/logrus"
)

type IRepository interface {
	Create(post *datastruct.Post) error
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

type IProducer interface {
	Produce(message []byte, userID uint64) error
}

type Service struct {
	repository    IRepository
	postCache     IPostCache
	friendService IFriendService
	producer      IProducer
}

func BuildService(db *postgres.DB, redisClient *redis.Client) *Service {
	p, err := rabbitmq.NewProducer()
	if err != nil {
		logger.Fatalf("Error creating producer: %v", err)
	}

	s := &Service{
		repository:    repository.NewRepository(db),
		friendService: friendService.BuildService(db),
		producer:      p,
	}

	s.postCache = cache.NewPostCacheJob(s, redisClient)
	return s
}

func (s *Service) Create(userID uint64, text string) error {
	post := &datastruct.Post{
		AuthorID:  userID,
		Text:      text,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := s.repository.Create(post); err != nil {
		logger.Errorf("Error creating post: %v", err)
		return err
	}
	s.SendNewPostToFriendsFeed(post, userID)
	return nil
}

func (s *Service) SendNewPostToFriendsFeed(post *datastruct.Post, userID uint64) {
	message, err := post.Encode()
	if err != nil {
		logger.Errorf("Error encoding post: %v", err)
		return
	}

	friends, err := s.friendService.GetUserFriends(userID)
	if err != nil {
		logger.Errorf("Error getting user's friends: %v", err)
		return
	}

	for _, friend := range friends {
		if !clients.IsConnectionForUser(friend) {
			continue
		}
		err = s.producer.Produce(message, friend)
		if err != nil {
			logger.Errorf("Error producing post: %v", err)
			return
		}
	}
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
			post := &datastruct.Post{
				AuthorID:  userID,
				Text:      scanner.Text(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			if err := s.repository.Create(post); err != nil {
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
