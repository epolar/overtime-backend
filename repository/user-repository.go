package repository

import (
	"orderStatistics/data"
)

type UserRepository struct {
	*Repository
}

func NewUserRepository(db *DBCli) *UserRepository {
	return &UserRepository{Repository: NewRepository(db)}
}

func DefaultUserRepository() *UserRepository {
	return NewUserRepository(DB())
}

func (r *UserRepository) SaveUser(user *data.User) error {
	return r.save(user)
}

func (r *UserRepository) GetAll() (list []*data.User, err error) {
	err = r.db.Find(&list).Error
	return
}

func (r *UserRepository) FindByID(userID uint64) (*data.User, error) {
	var user data.User
	if err := r.db.First(&user, "id = ?", userID).Error; err != nil {
		return nil, err
	} else {
		return &user, nil
	}
}

func (r *UserRepository) FindByIdIn(userIDs []uint64) (resp []*data.User, err error) {
	err = r.db.Find(&resp, "id in (?)", userIDs).Error
	return
}
