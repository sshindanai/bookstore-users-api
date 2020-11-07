package userscontroller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sshindanai/repo/bookstore-users-api/domain/usersdomain"
	"github.com/sshindanai/repo/bookstore-users-api/services/usersservices"
)

func CreateUser(c *gin.Context) {
	var user usersdomain.CreateUserRequest
	fmt.Println(user)
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		log.Fatal(err)
	}

	result, err := usersservices.UserService.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid parameters")
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
