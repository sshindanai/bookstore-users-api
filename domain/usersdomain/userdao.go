package usersdomain

import (
	"github.com/sshindanai/repo/bookstore-users-api/datasources/mysql/usersdb"
	"github.com/sshindanai/repo/bookstore-users-api/datasources/mysql/usersdbgorm"
	"github.com/sshindanai/repo/bookstore-users-api/utils/dateutils"
	"github.com/sshindanai/repo/bookstore-users-api/utils/errors"
	"github.com/sshindanai/repo/bookstore-users-api/utils/mysqlutils"
)

const (
	queryInsertUser       = "INSERT INTO users(first_name, last_name, email, date_created, date_updated, status, password) VALUES(?, ?, ?, ?, ?, ?, ?);"
	queryGetUsers         = "SELECT * FROM users;"
	queryGetUser          = "SELECT * FROM users WHERE ID=?;"
	queryUpdateUser       = "UPDATE users SET first_name=?, last_name=?, email=?, date_updated=? WHERE ID=?;"
	queryDeleteUser       = "DELETE FROM users WHERE ID=?;"
	queryFindUserByStatus = "SELECT * FROM users WHERE status=?;"
	statusActive          = "active"
)

func (user *User) Save() *errors.RestErr {
	// Prepare
	stmt, err := usersdb.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	// Execute
	insertResult, err := stmt.Exec(user.Firstname, user.Lastname, user.Email, user.DateCreated, user.DateUpdated, user.Status, user.Password)
	if err != nil {
		return mysqlutils.ParseError(err, user.Email)
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return mysqlutils.ParseError(err, userID)
	}
	user.ID = userID
	return nil
}

func (user *CreateUserRequest) CreateUser() (*User, *errors.RestErr) {
	newUser := &User{
		Firstname:   user.Firstname,
		Lastname:    user.Lastname,
		Email:       user.Email,
		DateCreated: dateutils.GetNowDBFormat(),
		DateUpdated: dateutils.GetNowDBFormat(),
		Status:      "active",
		Password:    user.Password,
	}

	if err := newUser.Save(); err != nil {
		return nil, err
	}
	return newUser, nil
}

func (user *CreateUserRequest) GormCreateUser() (*User, *errors.RestErr) {
	newUser := &User{
		Firstname:   user.Firstname,
		Lastname:    user.Lastname,
		Email:       user.Email,
		DateCreated: dateutils.GetNowDBFormat(),
		DateUpdated: dateutils.GetNowDBFormat(),
		Password:    user.Password,
	}

	result := usersdbgorm.GormDB.Create(&newUser)
	if result.Error != nil {
		return nil, errors.NewInternalServerError(result.Error.Error())
	}
	return newUser, nil
}

func (user *User) GetUser() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	// Scan is the mapping process
	if err := result.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email, &user.Status, &user.Password, &user.DateCreated, &user.DateUpdated); err != nil {
		return mysqlutils.ParseError(err, user.ID)
	}

	return nil
}

func (user *User) GormGetUser() *errors.RestErr {
	result := usersdbgorm.GormDB.First(&user)

	if result.Error != nil {
		return mysqlutils.ParseError(result.Error, user.ID)
	}
	return nil
}

func (user *User) UpdateUser() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Firstname, user.Lastname, user.Email, user.DateUpdated, user.ID)
	if err != nil {
		return mysqlutils.ParseError(err, user.Email)
	}

	return nil
}

func GetUserList() (*GetUsersDto, *errors.RestErr) {
	stmt, err := usersdb.Client.Prepare(queryGetUsers)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	var userList GetUsersDto
	result, _ := stmt.Query()
	for result.Next() {
		var user User
		if err := result.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email, &user.Status, &user.Password, &user.DateCreated, &user.DateUpdated); err != nil {
			return &userList, nil
		}
		userList.Users = append(userList.Users, user)
	}
	userList.TotalUsers = int64(len(userList.Users))
	return &userList, nil
}

func GormGetUserList() (*GetUsersDto, *errors.RestErr) {
	var users []User
	var userList GetUsersDto
	result := usersdbgorm.GormDB.Find(&users)
	if result.Error != nil {
		return nil, errors.NewInternalServerError(result.Error.Error())
	}
	userList.TotalUsers = result.RowsAffected
	userList.Users = users

	return &userList, nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, err = stmt.Exec(&user.ID); err != nil {
		return mysqlutils.ParseError(err, user.ID)
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := usersdb.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	var results = make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email, &user.Status, &user.Password, &user.DateCreated, &user.DateUpdated); err != nil {
			return results, nil
		}
		results = append(results, user)
	}

	return results, nil
}

func (user *User) FindByEmailAndPassword() *errors.RestErr {
	result := usersdbgorm.GormDB.Where("email = ? AND password = ? AND status = ?", user.Email, user.Password, statusActive).First(&user)
	if result.Error != nil {
		return errors.NewUnauthorizedError("invalid email or password")
	}
	return nil
}
