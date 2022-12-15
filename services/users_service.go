package services

import (
	"bookstore_users/domain/users"
	"bookstore_users/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	if userId <= 0 {
		return nil, errors.NewBadRequestError("invalid user id")
	}

	user := &users.User{Id: userId}
	if err := user.Get(); err != nil {
		return nil, err
	}

	return user, nil
}

func FindUser() {}
