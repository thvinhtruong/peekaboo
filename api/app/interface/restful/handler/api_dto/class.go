package api_dto

import (
	"server/utils/e"

	"github.com/gin-gonic/gin"
)

type ClassCreation struct {
	Class
}

func (cc ClassCreation) Validate(HasID bool, HasBody bool) error {
	if err := cc.Class.Validate(HasID, HasBody); err != nil {
		return err
	}

	return nil
}

func BindClassCreation(c *gin.Context, HasID bool, HasBody bool) (ClassCreation, error) {
	var d ClassCreation
	if err := c.ShouldBindJSON(&d); err != nil {
		return ClassCreation{}, e.ErrorBindJSON
	}

	if err := d.Validate(HasID, HasBody); err != nil {
		return ClassCreation{}, e.ErrorInputInvalid
	}

	return d, nil
}

type Class struct {
	ID           int    `json:"id"`
	Classname    string `json:"className"`
	Info         string `json:"info"`
	Announcement string `json:"announcement"`
	RoomCode     string `json:"roomCode"`
	Level        string `json:"level"`
}

func (c Class) Validate(HasID bool, HasBody bool) error {
	if HasID && c.ID == 0 {
		return e.ErrorInputInvalid
	}

	if HasBody {
		if err := CheckStringLength([]string{c.Classname, c.Info, c.Level, c.RoomCode}, 2, 100, false); err != nil {
			return e.ErrorInputInvalid
		}

		return nil
	}
	return nil
}

func BindClass(c *gin.Context, HasID bool, HasBody bool) (Class, error) {
	var nc Class
	if err := c.ShouldBindJSON(&nc); err != nil {
		return Class{}, e.ErrorBindJSON
	}

	if err := nc.Validate(HasID, HasBody); err != nil {
		return Class{}, e.ErrorInputInvalid
	}

	return nc, nil
}
