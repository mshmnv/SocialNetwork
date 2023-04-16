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
	AddUsers() error
	Search(firstName, secondName string) ([]datastruct.User, error)
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
		return status.Error(codes.NotFound, "User is not found")
	}

	dbData, err := s.repository.GetLoginData(data.ID)
	if err != nil {
		logger.Errorf("Error getting login data: %v", err)
		return err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(dbData.Password), []byte(data.Password)); err != nil {
		logger.Errorf("Wrong password: %v", err)
		return status.Error(codes.Unauthenticated, "Invalid password")
	}

	return nil
}

func (s *Service) Register(data *datastruct.User) error {
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

func (s *Service) GetUser(id uint64) (*datastruct.User, error) {

	if !s.IsExistedUser(id) {
		return nil, status.Error(codes.NotFound, "User is not found")
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

func (s *Service) Search(firstName, secondName string) ([]datastruct.User, error) {
	users, err := s.repository.Search(firstName, secondName)
	if err != nil {
		logger.Errorf("Error searching users: %v", err)
		return nil, err
	}
	return users, err
}

func (s *Service) AddUsers() {
	logger.Info("Process of adding users started")

	go func() {
		err := s.repository.AddUsers()
		if err != nil {
			logger.Errorf("Error adding users: %v", err)
			return
		}
		logger.Info("All users are successfully added")
	}()
}
