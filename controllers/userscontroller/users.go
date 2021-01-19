package userscontroller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sshindanai/repo/bookstore-users-api/domain/usersdomain"
	"github.com/sshindanai/repo/bookstore-users-api/services/usersservices"
	"github.com/sshindanai/repo/bookstore-users-api/utils/errors"
)

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	userId, idErr := strconv.ParseInt(userIdParam, 10, 64)
	if idErr != nil {
		return 0, errors.NewBadRequestError("id must be a number")
	}
	return userId, nil
}

func Create(c *gin.Context) {
	var user usersdomain.CreateUserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	result, err := usersservices.UserService.Create(&user)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func Get(c *gin.Context) {
	userId, err := getUserId(c.Param("id"))
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	result, err := usersservices.UserService.Get(userId)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, result.Marshal(c.GetHeader("X-Public") == "true"))
}

func GetUsers(c *gin.Context) {
	result, err := usersservices.UserService.GetUsers()

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, result.Marshal(c.GetHeader("X-Public") == "true"))
}

func FindUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not implement FindUser!")
}

func Update(c *gin.Context) {
	userId, err := getUserId(c.Param("id"))
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	var user usersdomain.CreateUserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	isPartial := c.Request.Method == http.MethodPatch

	result, err := usersservices.UserService.Update(isPartial, userId, &user)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, result)

}

func Delete(c *gin.Context) {
	userId, err := getUserId(c.Param("id"))
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	// Delete
	if err := usersservices.UserService.Delete(userId); err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": fmt.Sprintf("user %d is deleted", userId)})
}

func Search(c *gin.Context) {
	status := c.Query("status")

	users, err := usersservices.UserService.Search(status)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

func Login(c *gin.Context) {
	var request usersdomain.LoginUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Code, restErr)
		return
	}
	user, err := usersservices.UserService.LoginUser(&request)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, user)
}
