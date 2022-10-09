package users

import (
	"bookstore_users-api/datasources/mysql/users_db"
	"bookstore_users-api/util/customError"
	"bookstore_users-api/util/date_utils"
	"fmt"
	"log"
	"strings"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES (?,?,?,?);"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *customError.RestErr {
	log.Println("Sendign ping to db")
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	result := usersDB[user.Id]
	if result == nil {
		return customError.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *customError.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)

	if err != nil {
		return customError.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)

	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return customError.NewBadRequestError(fmt.Sprintf("Email %s already exists", user.Email))
		}

		return customError.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	userId, err := insertResult.LastInsertId()

	if err != nil {
		return customError.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	user.Id = userId

	return nil

	//current := usersDB[user.Id]
	//if current != nil {
	//	if current.Email == user.Email {
	//		return customError.NewBadRequestError(fmt.Sprintf("Email %s already registered", user.Id))
	//	}
	//	return customError.NewBadRequestError(fmt.Sprintf("User %d already exist", user.Id))
	//}
	//
	//user.DateCreated = date_utils.GetNowString()
	//
	//usersDB[user.Id] = user

	return nil
}
