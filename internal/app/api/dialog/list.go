package dialog

import (
	"context"

	"github.com/mshmnv/SocialNetwork/internal/app/service/dialog/datastruct"
	desc "github.com/mshmnv/SocialNetwork/pkg/api/dialog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (i Implementation) List(ctx context.Context, request *desc.ListRequest) (*desc.ListResponse, error) {
	senderID := ctx.Value("user_id").(uint64)
	if senderID == 0 {
		return nil, status.Error(codes.InvalidArgument, "Invalid sender id")
	}
	res, err := i.dialog.List(int64(senderID))
	if err != nil {
		return nil, err
	}

	return &desc.ListResponse{Messages: getMessages(res)}, nil
}

func getMessages(res []datastruct.Message) []*desc.ListResponse_DialogMessage {
	m := make([]*desc.ListResponse_DialogMessage, 0, len(res))
	for _, r := range res {
		m = append(m, &desc.ListResponse_DialogMessage{
			From:   r.Sender,
			To:     r.Receiver,
			Text:   r.Text,
			SentAt: timestamppb.New(r.SentAt),
		})
	}
	return m
}
