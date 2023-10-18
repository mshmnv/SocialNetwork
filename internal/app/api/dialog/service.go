package dialog

import (
	"github.com/mshmnv/SocialNetwork/internal/app/service/dialog/datastruct"
	"github.com/mshmnv/SocialNetwork/internal/app/service/dialog/service"
	"github.com/mshmnv/SocialNetwork/internal/pkg/postgres"
	dialogDesc "github.com/mshmnv/SocialNetwork/pkg/api/dialog"
)

type IDialogService interface {
	Send(sender, receiver int64, text string) error
	List(user int64) ([]datastruct.Message, error)
}

type Implementation struct {
	dialog IDialogService
	dialogDesc.UnimplementedDialogAPIServer
}

func NewDialogAPI(sharded *postgres.ShardedDB, db *postgres.DB) *Implementation {
	return &Implementation{
		dialog: service.BuildService(sharded, db),
	}
}
