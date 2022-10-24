package api_dto

import (
	"server/utils/e"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname"`
	Username string `json:"username"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
	// Address  string `json:"address"`
	Mail string `json:"mail"`
	// Phone         string `json:"phone"`
	// Dob           int64  `json:"dob"`
	// Qualification string `json:"qualification"`
	EntityCode int `json:"entity_code"`
}

type UserModel struct {
	ID         int    `json:"id"`
	EntityCode int    `json:"entityCode"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	FullName   string `json:"fullname"`
	Mail       string `json:"mail"`
	Gender     string `json:"gender"`
}

func BindUserModel(c *gin.Context, WithAuth bool) (UserModel, error) {
	var s UserModel
	if err := c.BindJSON(&s); err != nil {
		return UserModel{}, err
	}

	if WithAuth {
		if err := s.Validate(); err != nil {
			return UserModel{}, err
		}
	}

	return s, nil
}

func (a UserModel) Validate() error {
	if (len(a.Username) == 0) || (len(a.Password) == 0) {
		return e.ErrorInputInvalid
	}

	return nil
}

func BindAuthModel(c *gin.Context) (UserModel, error) {
	var a UserModel
	if err := c.ShouldBindJSON(&a); err != nil {
		return UserModel{}, err
	}

	if err := a.Validate(); err != nil {
		return UserModel{}, err
	}

	return a, nil
}
