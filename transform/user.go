package transform

import (
	"orderStatistics/data"
	"orderStatistics/response"
	"sync"
)

var userTransform *UserTransform
var userTransformOnce sync.Once

func User() *UserTransform {
	userTransformOnce.Do(func() {
		userTransform = &UserTransform{}
	})
	return userTransform
}

type UserTransform struct{}

func (c *UserTransform) ConvertUser(user *data.User) *response.User {
	return &response.User{
		Name: user.Name,
		ID:   user.ID,
	}
}

func (c *UserTransform) ConvertUserList(users []*data.User) []*response.User {
	list := make([]*response.User, 0, len(users))
	for _, user := range users {
		list = append(list, c.ConvertUser(user))
	}
	return list
}
