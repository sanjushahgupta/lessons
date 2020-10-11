package application

import (
	"github.com/sanjushahgupta/lessons/sample-rest-di/domain"
)

// CreateUser creates user
// validating if provided email exists
func (l *Logic) CreateUser(user domain.User) (int, error) {
	var id int

	users, err := l.store.GetAllUser()
	if err != nil {
		return id, err
	}

	for _, u := range users {
		if u.Email == user.Email {
			return id, domain.ErrEmailExist(user.Email)
		}
	}

	return l.store.CreateUser(user)
}

// GetAllUser gets all user
func (l *Logic) GetAllUser() ([]domain.User, error) {
	return l.store.GetAllUser()
}

// GetUser gets user
func (l *Logic) GetUser(id int) (domain.User, error) {
	return l.store.GetUser(id)
}
