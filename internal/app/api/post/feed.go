package post

import (
	"context"

	desc "github.com/mshmnv/SocialNetwork/pkg/api/post"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Feed реализует /post/feed
func (i *Implementation) Feed(ctx context.Context, req *desc.FeedRequest) (*desc.FeedResponse, error) {

	userID := ctx.Value("user_id").(uint64)
	if userID == 0 {
		return nil, status.Error(codes.InvalidArgument, "Invalid user id")
	}
	if req.GetOffset() == 0 && req.GetLimit() == 0 {
		req.Offset = 0
		req.Limit = 1000
	}

	feed, err := i.postService.Feed(userID, req.GetOffset(), req.GetLimit())
	if err != nil {
		return nil, err
	}

	result := &desc.FeedResponse{Feed: make([]*desc.Post, 0, len(feed))}

	for _, post := range feed {
		result.Feed = append(result.Feed, &desc.Post{
			Id:       post.PostID,
			Text:     post.Text,
			AuthorId: post.AuthorID,
		})
	}

	return result, nil
}
