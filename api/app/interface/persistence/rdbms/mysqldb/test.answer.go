package mysqldb

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"server/app/domain/entity"
)

func (q *Querier) CreateTestAnswer(ctx context.Context, ans entity.SubmittedAnswer) (err error) {
	sectionString, err := json.Marshal(&ans.Sections)
	if err != nil {
		return errors.New("failed to marshal section answer")
	}
	_, err = q.DB.ExecContext(ctx, "INSERT INTO test_answer (id, section_answer) VALUES (?, ?)", ans.ID, string(sectionString))
	return
}

func (q *Querier) DeleteTestAnswer(ctx context.Context, ans entity.SubmittedAnswer) (err error) {
	_, err = q.DB.ExecContext(ctx, "DELETE FROM test_answer WHERE id = ?", ans.ID)
	return
}

func (q *Querier) UpdateTestAnswer(ctx context.Context, ans entity.SubmittedAnswer) (err error) {
	sectionString, err := json.Marshal(&ans.Sections)
	if err != nil {
		return errors.New("failed to marshal section answer")
	}
	_, err = q.DB.ExecContext(ctx, "UPDATE test_answer SET section_answer = ? WHERE id = ?", string(sectionString), ans.ID)
	return
}

func (q *Querier) FindTestAnswer(ctx context.Context, id int) (ans entity.SubmittedAnswer, err error) {
	var sectionString string
	err = q.DB.QueryRowContext(ctx, "SELECT * FROM test_answer WHERE id = ?", id).Scan(&ans.ID, &sectionString)
	if err == sql.ErrNoRows {
		return entity.SubmittedAnswer{}, nil
	}

	if err := json.Unmarshal([]byte(sectionString), &ans.Sections); err != nil {
		return ans, errors.New("failed to unmarshal section answer")
	}

	return
}
