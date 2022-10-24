package api_dto

import (
	"server/utils/e"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var (
	Date = [...]string{"Mon", "Tues", "Weds", "Thu", "Fri", "Sat", "Sun"}
)

func ValidateDate(str string) error {
	for _, v := range Date {
		if str == v {
			return nil
		}
	}

	return e.ErrorWrongDateFormat
}

func CheckStringLength(strs []string, min int, max int, checkSpace bool) error {
	for _, str := range strs {
		err := validation.Validate(str, validation.Required, validation.Length(min, max), validation.When(checkSpace, validation.By(SpaceCheck)))
		if err == e.ErrorSpaceDetected {
			return err
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func SpaceCheck(value interface{}) error {
	s, _ := value.(string)
	if index := strings.Index(s, " "); index >= 0 && index <= len(s)-1 {
		return e.ErrorSpaceDetected
	}

	return nil
}
