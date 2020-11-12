package usersdomain

import (
	"strings"

	"github.com/sshindanai/repo/bookstore-users-api/utils/errors"
)

type CreateUserRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

type User struct {
	Id          int64  `json:"id"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	DateUpdated string `json:"date_updated"`
}

type GetUsersDto struct {
	Users      []User
	TotalUsers int64
}

func (user *CreateUserRequest) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	user.Firstname = strings.TrimSpace(strings.ToLower(user.Firstname))
	if user.Firstname == "" {
		return errors.NewBadRequestError("invalid firstname address")
	}
	user.Lastname = strings.TrimSpace(strings.ToLower(user.Lastname))
	if user.Lastname == "" {
		return errors.NewBadRequestError("invalid lastname address")
	}

	return nil
}
