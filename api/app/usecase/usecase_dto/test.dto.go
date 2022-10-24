package usecase_dto

type TestResult struct {
	ID          int    `json:"id"`
	TestID      int    `json:"testId"`
	TestName    string `json:"testName"`
	TestClassID int    `json:"testClassId"`
	UserID      int    `json:"userId"`
	EntityCode  int    `json:"entityCode"`
	DateCreated int64  `json:"dateCreated"`
	Score       int    `json:"score"`
	Comment     string `json:"comment"`
	ResultNote  string `json:"resultNote"`
}

type (
	Test struct {
		ID                   int    `json:"id"`
		TestClassID          int    `json:"test_class_id"`
		TagID                int    `json:"tag_id"`
		TestName             string `json:"test_name"`
		IsDone               bool   `json:"isDone"`
		PreviousTestResultID int    `json:"previousTestResultId"`
		CreatedUserID        int    `json:"created_user_id"`
		TargetEntityCode     int    `json:"target_entity_code"`
		Title                string `json:"title"`
		Info                 string `json:"info"`
		Duration             int    `json:"duration"`
		DateAssigned         int64  `json:"date_assigned"`
		Deadline             int64  `json:"deadline"`
		DateCreated          int64  `json:"dateCreated"`
		DateUpdated          int64  `json:"dateUpdated"`
	}

	SkillTest struct {
		ID          int       `db:"id"`
		MediaURL    string    `db:"mediaUrl"`
		Title       string    `db:"title"`
		Content     string    `db:"content"`
		Description string    `db:"description"`
		Type        string    `db:"type"`
		Section     []Section `db:"sections"`
	}

	Section struct {
		StartIndex int              `json:"startIndex"`
		EndIndex   int              `json:"endIndex"`
		Media      []SectionMedia   `json:"media"`
		Title      string           `json:"title"`
		Type       string           `json:"type"`
		Content    []SectionContent `json:"content"`

		DateCreated int64 `json:"dateCreated"`
		DateUpdated int64 `json:"dateUpdated"`
	}

	SectionContent struct {
		Q          string   `json:"q"`
		A          []string `json:"a"`
		CorrectAns string   `json:"correctAns"`
		ChosenAns  string   `json:"chosenAns"`
		// Explaination string   `json:"explaination"`
	}

	SectionMedia struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	TestClassRelation struct {
		ID      int `db:"id"`
		TestID  int `db:"test_id"`
		ClassID int `db:"class_id"`
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
