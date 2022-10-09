package controllers

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/services"
	"bookstore_users-api/util/customError"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := customError.NewBadRequestError("Invalid Json")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)

	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	param := c.Param("user_id")

	id, err := strconv.ParseInt(param, 10, 64)

	if err != nil {
		panic(err)
	}

	user, getUserError := services.GetUser(id)

	if getUserError != nil {
		c.JSON(getUserError.Status, getUserError)
	}

	c.JSON(http.StatusOK, user)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me")
}
