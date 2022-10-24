package api_dto

import (
	"server/utils/e"

	"github.com/gin-gonic/gin"
)

type (
	SkillTest struct {
		ID          int       `json:"id"`
		MediaURL    string    `json:"mediaURL"`
		Title       string    `json:"title"`
		Content     string    `json:"content"`
		Description string    `json:"description"`
		Type        string    `json:"type"`
		Section     []Section `json:"sections"`
	}

	Content struct {
		Q          string   `json:"q"`
		A          []string `json:"a"`
		CorrectAns string   `json:"correctAns"`
		ChosenAns  string   `json:"chosenAns"`
		// Explaination string   `json:"explaination"`
	}

	Media struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	Section struct {
		StartIndex             int       `json:"startIndex"`
		EndIndex               int       `json:"endIndex"`
		Title                  string    `json:"title"`
		Media                  []Media   `json:"media"`
		Type                   string    `json:"type"`
		SmallAnswerDescription string    `json:"smallAnswerDescription"`
		Content                []Content `json:"content"`
	}
	SubmitData struct {
		ID          int                 `json:"id"`
		TestClassID int                 `json:"testClassId"`
		Sections    []SubmitDataSection `json:"sections"`
	}

	SubmitDataSection struct {
		StartIndex int      `json:"startIndex"`
		EndIndex   int      `json:"endIndex"`
		Answers    []string `json:"answers"`
	}
)

func BindSubmitData(c *gin.Context) (SubmitData, error) {
	var s SubmitData
	err := c.BindJSON(&s)

	if s.ID == 0 {
		return s, e.ErrorInputInvalid
	}

	if len(s.Sections) == 0 {
		return s, e.ErrorInputInvalid
	}

	return s, err
}
