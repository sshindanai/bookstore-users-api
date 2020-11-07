package app

import (
	"github.com/sshindanai/repo/bookstore-users-api/controllers/userscontroller"
)

func mapURLs() {
	// Users endpoints
	router.GET("/users", userscontroller.GetUsers)
	router.GET("/users/:user_id", userscontroller.GetUser)
	//router.GET("/users/search", userscontroller.FindUser)
	router.POST("/users", userscontroller.CreateUser)
	router.PUT("/users/:user_id", userscontroller.UpdateUser)
	router.DELETE("/users/:user_id", userscontroller.DeleteUser)
}
