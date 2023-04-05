package service

import (
	"context"

	"github.com/mshmnv/SocialNetwork/internal/app/service/user/datastruct"
	"github.com/mshmnv/SocialNetwork/internal/app/service/user/repository"
	logger "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IRepository interface {
	Register(data *datastruct.User) error
	GetUser(id uint64) (*datastruct.User, error)
	IsExistedUser(id uint64) bool
	GetLoginData(id uint64) (*datastruct.LoginData, error)
}

type Service struct {
	repository IRepository
	ctx        context.Context
}

func BuildService(ctx context.Context) *Service {
	return &Service{
		repository: repository.NewRepository(ctx),
	}
}

func (s *Service) Login(data *datastruct.LoginData) error {

	if !s.IsExistedUser(data.ID) {
		return status.Error(codes.NotFound, "Пользователь не найден")
	}

	dbData, err := s.repository.GetLoginData(data.ID)
	if err != nil {
		logger.Errorf("Error getting login data: %v", err)
		return err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(dbData.Password), []byte(data.Password)); err != nil {
		logger.Errorf("Wrong password: %v", err)
		return status.Error(codes.Unauthenticated, "Неверный пароль")
	}

	return nil
}

func (s *Service) Register(ctx context.Context, data *datastruct.User) error {
	pass, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Errorf("Error generating user password: %v", err)
		return err
	}
	data.Password = string(pass)
	if err = s.repository.Register(data); err != nil {
		logger.Errorf("Error registering user: %v", err)
		return err
	}
	return nil
}

func (s *Service) GetUser(ctx context.Context, id uint64) (*datastruct.User, error) {

	if !s.IsExistedUser(id) {
		return nil, status.Error(codes.NotFound, "Анкета не найдена")
	}

	user, err := s.repository.GetUser(id)
	if err != nil {
		logger.Errorf("Error getting user data: %v", err)
		return nil, err
	}
	return user, nil
}

func (s *Service) IsExistedUser(id uint64) bool {
	return s.repository.IsExistedUser(id)
}
