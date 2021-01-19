package usersdomain

import "encoding/json"

// For presenting to the users
type PublicUser struct {
	ID          int64  `json:"id"`
	DateCreated string `json:"date_created" gorm:"column:date_created"`
	DateUpdated string `json:"date_updated" gorm:"column:date_updated"`
	Status      string `json:"status" gorm:"default:active"`
}

type PrivateUser struct {
	ID          int64  `json:"id"`
	Firstname   string `json:"firstname" gorm:"column:first_name"`
	Lastname    string `json:"lastname" gorm:"column:last_name"`
	Email       string `json:"email" gorm:"unique"`
	DateCreated string `json:"date_created" gorm:"column:date_created"`
	DateUpdated string `json:"date_updated" gorm:"column:date_updated"`
	Status      string `json:"status" gorm:"default:active"`
}

func (users GetUsersDto) Marshal(isPublic bool) interface{} {
	result := make([]interface{}, len(users.Users))
	for i, user := range users.Users {
		result[i] = user.Marshal(isPublic)
	}
	return result
}

func (user *User) Marshal(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			ID:          user.ID,
			DateCreated: user.DateCreated,
			DateUpdated: user.DateUpdated,
		}
	}
	userJson, _ := json.Marshal(user)
	var privateUser PrivateUser
	json.Unmarshal(userJson, &privateUser)
	return privateUser
}
