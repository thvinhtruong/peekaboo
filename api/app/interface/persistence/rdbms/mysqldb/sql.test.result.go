package mysqldb

import (
	"context"
	"database/sql"
	"server/app/domain/entity"
	"server/utils/conversion"
	"time"
)

var (
	sql_test_result_insert = "insert testresults SET test_class_id = ?, user_id = ?, entity_code = ?, datecreated = ?, score = ?, comment = ?, resultnote = ?, active = ?"

	sql_test_result_update = "update testresults SET score = ?, comment = ?, resultnote = ?, dateupdated = ? WHERE id = ? AND test_class_id = ? AND user_id = ? AND active = ?"
)

func insertTestResult(q *Querier, ctx context.Context, result entity.TestResult) (int, error) {
	date := conversion.ConvertUnixTimeMySqlTime(time.Now().Unix())
	r, err := q.DB.ExecContext(ctx, sql_test_result_insert, result.TestClassID, result.UserID, result.EntityCode, date, result.Score, result.Comment, result.ResultNote, 1)
	if err != nil {
		return 0, err
	}

	ID, err := r.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(ID), nil
}

func updateTestResult(q *Querier, ctx context.Context, result entity.TestResult) error {
	date := conversion.ConvertUnixTimeMySqlTime(time.Now().Unix())
	_, err := q.DB.ExecContext(ctx, sql_test_result_update, result.Score, result.Comment, result.ResultNote, date, result.ID, result.TestClassID, result.UserID, 1)
	if err != nil {
		return err
	}

	return nil
}

// Do not select the Active, DateCreated and DateUpdated.
func refactorTestResultDetail(rows *sql.Rows) (data []entity.TestResult, err error) {
	for rows.Next() {
		var created string
		var updated, comment, note sql.NullString
		var result entity.TestResult
		var score sql.NullInt64

		if err := rows.Scan(&result.ID, &result.TestClassID, &result.UserID, &result.EntityCode, &created, &score, &comment, &note, &result.Active, &updated); err != nil {
			return data, err
		}

		if score.Valid {
			result.Score = int(score.Int64)
		}

		if updated.Valid {
			result.DateUpdated = conversion.ConvertMysqlTimeUnixTime(updated.String)
		}

		if note.Valid {
			result.ResultNote = note.String
		}

		if comment.Valid {
			result.Comment = comment.String
		}

		result.DateCreated = conversion.ConvertMysqlTimeUnixTime(created)

		data = append(data, result)
	}

	return data, nil
}

// Do not select the Active, DateCreated and DateUpdated.
func refactorTestResultIndex(rows *sql.Rows) (data []entity.TestResult, err error) {
	for rows.Next() {
		var created string
		var score sql.NullFloat64
		var result entity.TestResult
		var comment, resultNote sql.NullString

		if err := rows.Scan(&result.ID, &result.TestClassID, &result.UserID, &result.EntityCode, &created, &score, &comment, &resultNote); err != nil {
			return data, err
		}

		if score.Valid {
			result.Score = int(score.Float64)
		}

		if comment.Valid {
			result.Comment = comment.String
		}

		if resultNote.Valid {
			result.ResultNote = resultNote.String
		}

		result.DateCreated = conversion.ConvertMysqlTimeUnixTime(created)

		data = append(data, result)
	}

	return data, nil
}
