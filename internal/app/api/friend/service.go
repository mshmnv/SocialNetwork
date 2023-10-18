package friend

import (
	friendService "github.com/mshmnv/SocialNetwork/internal/app/service/friend/service"
	"github.com/mshmnv/SocialNetwork/internal/pkg/postgres"
	friendAPI "github.com/mshmnv/SocialNetwork/pkg/api/friend"
)

type IFriendService interface {
	SendFriendRequest(friendID uint64, userID uint64) error
	DeleteFriend(friendID uint64, userID uint64) error
}

type Implementation struct {
	friendService IFriendService
	friendAPI.UnimplementedFriendAPIServer
}

func NewFriendAPI(db *postgres.DB) *Implementation {
	return &Implementation{
		friendService: friendService.BuildService(db),
	}
}
