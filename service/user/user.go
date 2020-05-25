package user

import (
	"errors"
	"orderStatistics/data"
	"orderStatistics/repository"
	"orderStatistics/runtime/log"
	"sync"

	"github.com/jinzhu/gorm"
)

var userService User
var userServiceOnce sync.Once

func Service() *User {
	userServiceOnce.Do(func() {
		userService = User{}
	})
	return &userService
}

type User struct{}

func (u *User) Add(
	name string,
	label string,
	nick string,
) (*data.User, error) {
	rep := repository.DefaultUserRepository()

	user := &data.User{
		Name:  name,
		Label: label,
		Nick:  nick,
	}
	if err := rep.SaveUser(user); err != nil {
		log.Log.Errorf("save user info failure: %s", err)
		return nil, err
	}

	return user, nil
}

func (u *User) All() ([]*data.User, error) {
	rep := repository.DefaultUserRepository()
	return rep.GetAll()
}

func (u *User) GetUserInfoByID(userID uint64) (*data.User, error) {
	rep := repository.DefaultUserRepository()
	return rep.FindByID(userID)
}

func (u *User) Delete(userID uint64) (err error) {
	rep := repository.DefaultUserRepository()
	var user *data.User
	if user, err = rep.FindByID(userID); err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = errors.New("user not found")
		} else {
			log.Log.Errorf("")
		}
		return
	}
	return rep.DeleteUser(user)
}
