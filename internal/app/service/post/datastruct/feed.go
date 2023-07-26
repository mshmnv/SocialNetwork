package datastruct

import (
	"sync"
)

type user uint64

type UserPosts struct {
	sync.RWMutex
	Posts map[user][]uint64
}

func NewUserPosts() UserPosts {
	return UserPosts{
		Posts: make(map[user][]uint64),
	}
}

func (u *UserPosts) GetPostsForUsers(userIDs []uint64) []uint64 {
	u.RLock()
	defer u.RUnlock()
	result := make([]uint64, 0)
	for _, userID := range userIDs {
		if posts, ok := u.Posts[user(userID)]; ok {
			result = append(result, posts...)
		}
	}
	return result
}

// AddUserPost добавить новый пост пользователя
func (u *UserPosts) AddUserPost(userID uint64, postID uint64) {
	u.Lock()
	defer u.Unlock()
	u.Posts[user(userID)] = appendIfNotExists(u.Posts[user(userID)], postID)
}

// DeleteUserPost удалить пост пользователя
func (u *UserPosts) DeleteUserPost(userID, postID uint64) {
	oldPosts := u.Posts[user(userID)]
	newPosts := make([]uint64, 0, len(oldPosts))

	for _, post := range u.Posts[user(userID)] {
		if post == postID {
			continue
		}
		newPosts = append(newPosts, post)
	}

	u.Lock()
	defer u.Unlock()
	u.Posts[user(userID)] = newPosts
}

func appendIfNotExists(posts []uint64, postID uint64) []uint64 {
	for _, post := range posts {
		if post == postID {
			return posts
		}
	}
	return append(posts, postID)
}
