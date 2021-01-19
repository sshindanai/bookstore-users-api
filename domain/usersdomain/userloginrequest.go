package usersdomain

import (
	"strings"

	"github.com/sshindanai/repo/bookstore-users-api/utils/errors"
)

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l *LoginUserRequest) Validate() *errors.RestErr {
	if strings.TrimSpace(l.Email) == "" {
		return errors.NewBadRequestError("empty email is invalid")
	}
	if strings.TrimSpace(l.Password) == "" {
		return errors.NewBadRequestError("empty password is invalid")
	}
	return nil
}
