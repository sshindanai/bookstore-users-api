package usersservices

import (
	"github.com/sshindanai/repo/bookstore-users-api/domain/usersdomain"
	"github.com/sshindanai/repo/bookstore-users-api/utils/errors"
)

type userService struct{}

type userInterface interface {
	CreateUser(*usersdomain.CreateUserRequest) (*usersdomain.User, *errors.RestErr)
	//GetUsers() (*usersdomain.GetUsersDto, error)
	//GetUser(id int64) (*usersdomain.User, error)
}

var (
	UserService userInterface
)

func init() {
	UserService = &userService{}
}

func (s *userService) CreateUser(user *usersdomain.CreateUserRequest) (*usersdomain.User, *errors.RestErr) {
	// Validate inputs
	if err := user.Validate(); err != nil {
		return nil, err
	}

	// Return mock result
	result, err := user.CreateUser()
	if err != nil {
		return nil, err
	}

	return result, nil
}
