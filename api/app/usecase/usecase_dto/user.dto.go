package usecase_dto

type ReviewTestUserRequest struct {
	ID      int    `json:"test_result_id"`
	UserID  int    `json:"user_id"`
	Score   int    `json:"score"`
	Comment string `json:"comment"`
}

type User struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname"`
	Username string `json:"username"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
	Mail     string `json:"mail"`
	// Phone         string `json:"phone"`
	// Dob           int64  `json:"dob"`
	// Qualification string `json:"qualification"`
	EntityCode int `json:"entity_code"`
}
