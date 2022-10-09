package services

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/util/customError"
)

func CreateUser(user users.User) (*users.User, *customError.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(id int64) (*users.User, *customError.RestErr) {
	user := &users.User{
		Id: id,
	}

	err := user.Get()

	if err != nil {
		return nil, err
	}

	return user, nil
}
