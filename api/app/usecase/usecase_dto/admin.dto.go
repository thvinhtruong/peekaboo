package usecase_dto

type UpdatedTestResultResponse struct {
	ID      int    `json:"test_result_id"`
	UserID  int    `json:"user_id"`
	Score   int    `json:"score"`
	Comment string `json:"comment"`
}

type DeletedTestResultResponse struct {
	ID int `json:"test_result_id"`
}

type ResponseUserReviewRequest struct {
	ID     int `json:"test_result_id"`
	UserID int `json:"user_id"`
	Score  int `json:"score"`
}
