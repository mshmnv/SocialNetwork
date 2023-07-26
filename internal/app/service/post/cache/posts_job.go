package cache

import (
	"time"

	"github.com/mshmnv/SocialNetwork/internal/app/service/post/datastruct"
	logger "github.com/sirupsen/logrus"
)

type IPostService interface {
	GetPostsAfterDate(date time.Time) ([]datastruct.Post, error)
}

type IPostCache interface {
	UpdatePost(post *datastruct.Post)
}

type PostUpdateJob struct {
	lastUpdateTime time.Time
	postService    IPostService
	cache          IPostCache
}

func StartPostUpdateJob(cache IPostCache, postService IPostService) {
	p := &PostUpdateJob{
		lastUpdateTime: time.Now().Add(-expiration),
		cache:          cache,
		postService:    postService,
	}
	p.Exec()
}

func (p *PostUpdateJob) Exec() {
	for {
		posts, err := p.postService.GetPostsAfterDate(p.lastUpdateTime)
		if err != nil {
			logger.Errorf("Error getting posts updates: %v", err)
		}

		// update cache
		for _, post := range posts {
			p.cache.UpdatePost(&post)
		}

		p.calculateLastUpdate(posts)

		time.Sleep(time.Minute)
	}
}

func (p *PostUpdateJob) calculateLastUpdate(posts []datastruct.Post) {
	for _, post := range posts {
		if post.UpdatedAt.After(p.lastUpdateTime) {
			p.lastUpdateTime = post.UpdatedAt
		}
	}
}
