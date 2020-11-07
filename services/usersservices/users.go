package usersservices

import (
	"crypto/rand"
	"fmt"
	"log"
	"time"

	"github.com/sshindanai/repo/bookstore-users-api/domain/usersdomain"
)

type userService struct{}

type userInterface interface {
	CreateUser(*usersdomain.CreateUserRequest) (*usersdomain.User, error)
	//GetUsers() (*usersdomain.GetUsersDto, error)
	//GetUser(id int64) (*usersdomain.User, error)
}

var (
	UserService userInterface
)

func init() {
	UserService = &userService{}
}

func (s *userService) CreateUser(request *usersdomain.CreateUserRequest) (*usersdomain.User, error) {
	// Validate inputs
	// Return mock result
	mockResult := &usersdomain.User{
		Id:          generateUUID(),
		Firstname:   request.Firstname,
		Lastname:    request.Lastname,
		Email:       request.Email,
		DateCreated: time.Now().Format(time.Stamp),
		DateUpdated: time.Now().Format(time.Stamp),
	}

	return mockResult, nil
}

func generateUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return uuid
}
