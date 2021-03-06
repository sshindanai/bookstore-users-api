package usersservices

import (
	"github.com/sshindanai/repo/bookstore-users-api/domain/usersdomain"
	"github.com/sshindanai/repo/bookstore-users-api/utils/cryptoutils"
	"github.com/sshindanai/repo/bookstore-users-api/utils/dateutils"
	"github.com/sshindanai/repo/bookstore-users-api/utils/errors"
)

// Make userService struct staisfy the userServiceInterface interface that name UserService
var (
	UserService userServiceInterface = &userService{}
)

type userService struct{}

type userServiceInterface interface {
	Create(*usersdomain.CreateUserRequest) (*usersdomain.User, *errors.RestErr)
	Get(int64) (*usersdomain.User, *errors.RestErr)
	GetUsers() (*usersdomain.GetUsersDto, *errors.RestErr)
	Update(bool, int64, *usersdomain.CreateUserRequest) (*usersdomain.User, *errors.RestErr)
	Delete(int64) *errors.RestErr
	Search(string) ([]usersdomain.User, *errors.RestErr)
	LoginUser(*usersdomain.LoginUserRequest) (*usersdomain.User, *errors.RestErr)
}

func (s *userService) Create(user *usersdomain.CreateUserRequest) (*usersdomain.User, *errors.RestErr) {
	// ValIDate inputs
	if err := user.Validate(); err != nil {
		return nil, err
	}
	// Hashing password
	user.Password = cryptoutils.GetSHA256(user.Password)

	result, err := user.GormCreateUser()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *userService) Get(userID int64) (*usersdomain.User, *errors.RestErr) {

	result := &usersdomain.User{ID: userID}
	if err := result.GormGetUser(); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *userService) GetUsers() (*usersdomain.GetUsersDto, *errors.RestErr) {
	result, err := usersdomain.GormGetUserList()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *userService) Update(isPartial bool, userID int64, user *usersdomain.CreateUserRequest) (*usersdomain.User, *errors.RestErr) {
	// Get current user by ID
	currUser, err := UserService.Get(userID)
	if err != nil {
		return nil, err
	}

	// Update
	if isPartial {
		if user.Firstname != "" {
			currUser.Firstname = user.Firstname
		}
		if user.Lastname != "" {
			currUser.Lastname = user.Lastname
		}
		if user.Email != "" {
			currUser.Email = user.Email
		}
	} else {
		if err := user.Validate(); err != nil {
			return nil, err
		}
		currUser.Firstname = user.Firstname
		currUser.Lastname = user.Lastname
		currUser.Email = user.Email
	}

	currUser.DateUpdated = dateutils.GetNowDBFormat()

	if err := currUser.UpdateUser(); err != nil {
		return nil, err
	}

	return currUser, nil
}

func (s *userService) Delete(userID int64) *errors.RestErr {
	// ValIDate check ID existing
	_, err := UserService.Get(userID)
	if err != nil {
		return err
	}

	// Delete
	user := &usersdomain.User{ID: userID}
	return user.Delete()
}

func (s *userService) Search(status string) ([]usersdomain.User, *errors.RestErr) {
	dao := &usersdomain.User{}
	users, err := dao.FindByStatus(status)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *userService) LoginUser(request *usersdomain.LoginUserRequest) (*usersdomain.User, *errors.RestErr) {
	err := request.Validate()
	if err != nil {
		return nil, err
	}

	user := &usersdomain.User{
		Email:    request.Email,
		Password: cryptoutils.GetSHA256(request.Password),
	}
	if err := user.FindByEmailAndPassword(); err != nil {
		return nil, err
	}

	// Authenticate
	// result, err := oauth.NewRestOauth().Authenticate(user)
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Println(result)

	return user, nil
}
