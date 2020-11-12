package userscontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sshindanai/repo/bookstore-users-api/domain/usersdomain"
	"github.com/sshindanai/repo/bookstore-users-api/services/usersservices"
	"github.com/sshindanai/repo/bookstore-users-api/utils/errors"
)

func CreateUser(c *gin.Context) {
	var user usersdomain.CreateUserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	result, err := usersservices.UserService.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUsers(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not implement GetUsers!")
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not implement GetUser!")

}

func FindUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not implement FindUser!")

}

func UpdateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not implement UpdateUser!")

}

func DeleteUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not implement DeleteUser!")
}
