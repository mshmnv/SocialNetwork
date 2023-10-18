package service

import (
	"github.com/mshmnv/SocialNetwork/internal/app/service/friend/repository"
	"github.com/mshmnv/SocialNetwork/internal/pkg/postgres"
)

type IRepository interface {
	SendFriendRequest(friendID uint64, userID uint64) error
	ApproveFriendRequest(friendID uint64, userID uint64) error
	DeleteFriend(friendID uint64, userID uint64) error
	GetUserFriends(userID uint64) ([]uint64, error)
}

type Service struct {
	repository IRepository
}

func BuildService(db *postgres.DB) *Service {
	return &Service{
		repository: repository.NewRepository(db),
	}
}

func (s *Service) SendFriendRequest(friendID uint64, userID uint64) error {
	return s.repository.SendFriendRequest(friendID, userID)
}

func (s *Service) ApproveFriendRequest(friendID uint64, userID uint64) error {
	return s.repository.ApproveFriendRequest(friendID, userID)
}

func (s *Service) DeleteFriend(friendID uint64, userID uint64) error {
	return s.repository.DeleteFriend(friendID, userID)
}

func (s *Service) GetUserFriends(userID uint64) ([]uint64, error) {
	return s.repository.GetUserFriends(userID)
}
