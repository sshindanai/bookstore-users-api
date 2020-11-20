package app

import (
	"github.com/sshindanai/repo/bookstore-users-api/controllers/userscontroller"
)

func mapURLs() {
	// Users endpoints
	router.GET("/users", userscontroller.GetUsers)
	router.GET("/users/:id", userscontroller.Get)
	router.GET("/internal/users/search", userscontroller.Search)
	router.POST("/users", userscontroller.Create)
	router.PUT("/users/:id", userscontroller.Update)
	router.PATCH("/users/:id", userscontroller.Update)
	router.DELETE("/users/:id", userscontroller.Delete)
}
