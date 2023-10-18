package dialog

import (
	"context"

	desc "github.com/mshmnv/SocialNetwork/pkg/api/dialog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i Implementation) Send(ctx context.Context, request *desc.SendRequest) (*desc.SendResponse, error) {
	senderID := ctx.Value("user_id").(uint64)
	if senderID == 0 {
		return nil, status.Error(codes.InvalidArgument, "Invalid sender id")
	}
	err := i.dialog.Send(int64(senderID), int64(request.GetUserId()), request.GetText())
	if err != nil {
		return nil, err
	}

	return &desc.SendResponse{}, nil
}
