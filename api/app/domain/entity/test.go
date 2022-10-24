package entity

type (
	//Test is the test bank. There can be unlimited tests.
	//TestName is the Topic/Assigned Name, etc.. of the test.
	//TestName is used for init Query, therefore it should be carefully cared by trimming both ends.Content is the question itself.
	//Each test will have a unique ID.
	//Active means that does the PTest still exist [1] or not [0].
	//Tag is based on the test_tag table, which is addable.
	//Difficulty base on Linkert scales: 1 is smallest - 5 is largest.
	//Duration is the time limit for the test.
	Test struct {
		ID               int    `db:"id"`
		TagID            int    `db:"tag_id"`
		TestName         string `db:"test_name"`
		CreatedUserID    int    `db:"created_user_id"`
		TargetEntityCode int    `db:"target_entity_code"`
		Title            string `db:"title"`
		Info             string `db:"info"`
		Duration         int    `db:"duration"`
		DateAssigned     int64  `db:"date_assigned"`
		Deadline         int64  `db:"deadline"`

		Active      int   `db:"active"`
		DateCreated int64 `db:"datecreated"`
		DateUpdated int64 `db:"dateupdated"`
		// Date, when inserting to database, should be set to type "yyyy-mm-dd"
	}

	SkillTest struct {
		Id          int       `db:"id"`
		MediaURL    string    `db:"media_url"`
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
		CorrectAns string   `json:"correct_ans"`
		ChosenAns  string   `json:"chosen_ans"`
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

	SubmittedAnswer struct {
		ID       int                      `json:"id"`
		Sections []SubmittedAnswerSection `json:"submitted_sections"`
	}

	SubmittedAnswerSection struct {
		StartIndex int      `json:"startIndex"`
		EndIndex   int      `json:"endIndex"`
		Answers    []string `json:"answers"`
	}
)
