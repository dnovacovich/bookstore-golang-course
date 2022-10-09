package users

import (
	"bookstore_users-api/util/customError"
	"strings"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	DateCreated string `json:"dateCreated"`
}

func (user *User) Validate() *customError.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return customError.NewBadRequestError("Invalid email address")
	}

	return nil
}
