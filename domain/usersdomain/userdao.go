package usersdomain

import (
	"fmt"

	"github.com/sshindanai/repo/bookstore-users-api/datasources/mysql/usersdb"
	"github.com/sshindanai/repo/bookstore-users-api/utils/dateutils"
	"github.com/sshindanai/repo/bookstore-users-api/utils/errors"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created, date_updated) VALUES(?, ?, ?, ?, ?);"
)

func (user *User) Save() *errors.RestErr {
	// Prepare
	statement, err := usersdb.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer statement.Close()

	// Execute
	insertResult, err := statement.Exec(user.Firstname, user.Lastname, user.Email, user.DateCreated, user.DateUpdated)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}
	user.Id = userId
	return nil
}

func (user *CreateUserRequest) CreateUser() (*User, *errors.RestErr) {
	if err := usersdb.Client.Ping(); err != nil {
		panic(err)
	}

	newUser := &User{
		Firstname:   user.Firstname,
		Lastname:    user.Lastname,
		Email:       user.Email,
		DateCreated: dateutils.GetNowString(),
		DateUpdated: dateutils.GetNowString(),
	}

	if err := newUser.Save(); err != nil {
		return nil, err
	}
	return newUser, nil
}

// func Get(userId int64) (*User, *errors.RestErr) {
// 	result := userDB[userId]
// 	if result == nil {
// 		return nil, errors.NewNotFoundError(fmt.Sprintf("user %s not found", userId))
// 	}

// 	result.Save()

// 	return result, nil
// }
