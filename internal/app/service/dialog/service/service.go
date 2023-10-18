package service

import (
	"sort"

	"github.com/mshmnv/SocialNetwork/internal/app/service/dialog/datastruct"
	"github.com/mshmnv/SocialNetwork/internal/app/service/dialog/repository"
	"github.com/mshmnv/SocialNetwork/internal/pkg/postgres"
	logger "github.com/sirupsen/logrus"
)

type IRepository interface {
	GetDialogID(user1, user2 int64) (int64, error)
	Add(dialogID, sender, receiver int64, text string) error

	GetDialogs(user int64) ([]int64, error)
	GetMessages(dialogIDs []int64) ([]datastruct.Message, error)
}

type Service struct {
	repository IRepository
}

func BuildService(sharded *postgres.ShardedDB, db *postgres.DB) *Service {
	return &Service{
		repository: repository.NewRepository(sharded, db),
	}
}

func (s *Service) Send(sender, receiver int64, text string) error {
	dialogID, err := s.repository.GetDialogID(sender, receiver)
	if err != nil || dialogID == 0 {
		logger.Errorf("Error getting dialog id for users - sender %d; receiver: %d; %s", sender, receiver, err)
		return err
	}

	if err := s.repository.Add(dialogID, sender, receiver, text); err != nil {
		logger.Errorf("Error sending user messages: %s", err)
		return err
	}
	return nil
}

func (s *Service) List(user int64) ([]datastruct.Message, error) {
	dialogs, err := s.repository.GetDialogs(user)
	if err != nil {
		logger.Errorf("Error getting list of user dialogs: %s", err)
		return nil, err
	}
	if len(dialogs) == 0 {
		logger.Infof("No dialogs for user %d", user)
	}

	res, err := s.repository.GetMessages(dialogs)
	if err != nil {
		logger.Errorf("Error getting list of user messages: %s", err)
		return nil, err
	}

	sort.Slice(res, func(i, j int) bool {
		if res[i].DialogID == res[j].DialogID {
			return res[i].SentAt.Before(res[j].SentAt)
		}
		return res[i].DialogID < res[j].DialogID
	})
	return res, nil
}
