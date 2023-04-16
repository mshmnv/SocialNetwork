package user

import (
	"context"

	desc "github.com/mshmnv/SocialNetwork/pkg/api/user"
)

// Search реализует /user/search
func (i *Implementation) Search(ctx context.Context, req *desc.SearchRequest) (*desc.SearchResponse, error) {

	users, err := i.userService.Search(req.GetFirstName(), req.GetSecondName())
	if err != nil {
		return nil, err
	}

	var userResults []*desc.UserData
	for _, user := range users {
		userResults = append(userResults, &desc.UserData{
			FirstName:  user.FirstName,
			SecondName: user.SecondName,
			Age:        user.Age,
			Birthdate:  user.BirthDate,
			Biography:  user.Biography,
			City:       user.City,
		})
	}

	return &desc.SearchResponse{Users: userResults}, nil
}
