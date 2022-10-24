package api_dto

import (
	"server/utils/e"

	"github.com/gin-gonic/gin"
)

type Test struct {
	ID                   int    `json:"id"`
	TestClassID          int    `json:"testClassId"`
	TagID                int    `json:"tagId"`
	TagName              string `json:"tagName"`
	TestName             string `json:"testName"`
	CreatedUserID        int    `json:"createdUserId"`
	PreviousTestResultID int    `json:"previousTestResultId"`
	TargetEntityCode     int    `json:"targetEntityCode"`
	Title                string `json:"title"`
	Info                 string `json:"info"`
	Duration             int    `json:"duration"`
	Status               string `json:"status"`
	IsDone               bool   `json:"isDone"`
	DateAssigned         int64  `json:"dateAssigned"`
	Deadline             int64  `json:"deadline"`
	DateCreated          int64  `json:"dateCreated"`
	DateUpdated          int64  `json:"dateUpdated"`
}

func (t Test) Validate(WithID bool, WithBody bool) error {
	if (WithID) && (t.ID == 0) {
		return e.ErrorInputInvalid
	}

	if WithBody {
		if (len(t.TestName) == 0) || (len(t.Info) == 0) || (t.DateAssigned == 0) || (t.Deadline == 0) {
			return e.ErrorInputInvalid
		}

		if t.Duration == 0 {
			return e.ErrorInputInvalid
		}
	}

	return nil
}

func BindTest(c *gin.Context, WithID bool, WithBody bool) (Test, error) {
	var t Test
	if err := c.ShouldBindJSON(&t); err != nil {
		return t, err
	}

	if err := t.Validate(WithID, WithBody); err != nil {
		return t, err
	}

	return t, nil
}
