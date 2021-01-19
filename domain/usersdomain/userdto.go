package usersdomain

import (
	"strings"

	"github.com/sshindanai/repo/bookstore-users-api/utils/errors"
)

type CreateUserRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type User struct {
	ID          int64  `json:"id"`
	Firstname   string `json:"firstname" gorm:"column:first_name"`
	Lastname    string `json:"lastname" gorm:"column:last_name"`
	Email       string `json:"email" gorm:"unique"`
	DateCreated string `json:"date_created" gorm:"column:date_created"`
	DateUpdated string `json:"date_updated" gorm:"column:date_updated"`
	Status      string `json:"status" gorm:"default:active"`
	Password    string `json:"password"`
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
		return errors.NewBadRequestError("invalid firstname")
	}
	user.Lastname = strings.TrimSpace(strings.ToLower(user.Lastname))
	if user.Lastname == "" {
		return errors.NewBadRequestError("invalid lastname")
	}
	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}

	return nil
}
