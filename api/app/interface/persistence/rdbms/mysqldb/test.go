package mysqldb

import (
	"context"
	"database/sql"
	"errors"
	"server/app/domain/entity"
	"server/utils/conversion"
	"time"
)

//This function do not adding question to test. Its purpose is to create a test record with testID.
func (q *Querier) CreateTest(ctx context.Context, test entity.Test) (int, error) {
	stmt, err := q.DB.PrepareContext(ctx, "INSERT testbank SET tag_id = ?, test_name = ?, created_user_id = ?, target_entity_code = ?, title = ?, info = ?, duration = ?, dateassigned = ?, deadline = ?, active = ?, datecreated = ?")
	if err != nil {
		return 0, err
	}

	date := conversion.ConvertUnixTimeMySqlTime(time.Now().Unix())
	dateassigned := conversion.ConvertUnixTimeMySqlTime(time.Now().Unix())
	deadline := conversion.ConvertUnixTimeMySqlTime(time.Now().Unix())

	result, err := stmt.ExecContext(ctx, test.TagID, test.TestName, test.CreatedUserID, test.TargetEntityCode, test.Title, test.Info, test.Duration, dateassigned, deadline, 1, date)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	k := int(id)

	return k, nil
}

//TestID will be considered first if a not-nil value is passed, else the bookname will be considered.
func (q *Querier) QueryTestHeadline(ctx context.Context, ID int, TestName string) (result []entity.Test, err error) {
	if (ID == 0) && (len(TestName) == 0) {
		return result, errors.New("invalid Input")
	}

	var data *sql.Rows
	if ID != 0 {
		data, err = q.DB.QueryContext(ctx, "Select * from testbank where id = ? and active = ?", ID, 1)
	} else {
		data, err = q.DB.QueryContext(ctx, "Select * from testbank where test_name = ? and active = ?", TestName, 1)
	}

	tempt, err := refactorQueryTest(data, err)
	if err != nil {
		return result, err
	}

	return tempt, nil
}

func (q *Querier) QueryAllTest(ctx context.Context) (result []entity.Test, err error) {
	data, err := q.DB.QueryContext(ctx, "Select * from testbank where active = ?", 1)
	if err == sql.ErrNoRows {
		return result, nil
	}

	if err != nil {
		return result, err
	}

	tempt, err := refactorQueryTest(data, err)
	if err != nil {
		return result, err
	}

	return tempt, nil
}

func (q *Querier) QueryTestByTestTag(ctx context.Context, TestTagID int) (result []entity.Test, err error) {
	if TestTagID == 0 {
		return result, errors.New("invalid input")
	}

	rows, err := q.DB.QueryContext(ctx, "Select * from testbank where tag_id = ? and active = ?", TestTagID, 1)
	if err != nil {
		return result, err
	}

	return refactorQueryTest(rows, nil)
}

func (q *Querier) QueryAllTestHeadlines(ctx context.Context) ([]entity.Test, error) {
	rows, err := q.DB.QueryContext(ctx, "SELECT * FROM testbank where active = ?", 1)

	if err == sql.ErrNoRows {
		return []entity.Test{}, nil
	}

	if err != nil {
		return []entity.Test{}, err
	}

	return refactorQueryTest(rows, nil)
}

// Replaceable fields: tagid, testname, info, ispublished, deadline
func (q *Querier) UpdateTest(ctx context.Context, test entity.Test) error {
	stmt, err := q.DB.PrepareContext(ctx, "UPDATE testbank SET tag_id = ?, test_name = ?, info = ?, target_entity_code = ?, title = ?, info = ?, duration = ?, deadline = ?, dateupdated = ? WHERE id = ? and active = ?")
	if err != nil {
		return err
	}

	date := conversion.ConvertUnixTimeMySqlTime(time.Now().Unix())
	deadline := conversion.ConvertUnixTimeMySqlTime(time.Now().Unix())

	_, err = stmt.ExecContext(ctx, test.TagID, test.TestName, test.Info, test.TargetEntityCode, test.Title, test.Info, test.Duration, deadline, date, test.ID, 1)
	if err != nil {
		return err
	}

	return nil
}

func (q *Querier) DeleteTest(ctx context.Context, TestID int) error {
	stmt, err := q.DB.PrepareContext(ctx, "Update testbank set active = ? where id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, 0, TestID)
	if err != nil {
		return err
	}

	return nil
}

func refactorQueryTest(r *sql.Rows, err error) ([]entity.Test, error) {
	var finalresult []entity.Test

	if err != nil {
		return []entity.Test{}, err
	}

	for r.Next() {
		var test entity.Test

		var duration sql.NullInt32
		var created string
		var assigned string
		var deadline sql.NullString
		var nullable sql.NullString

		err = r.Scan(&test.ID, &test.TagID, &test.TestName, &test.CreatedUserID, &test.TargetEntityCode, &test.Title, &test.Info, &test.Duration,
			&assigned, &deadline, &test.Active, &created, &nullable)

		if err != nil {
			return []entity.Test{}, err
		}

		if nullable.Valid {
			test.DateUpdated = conversion.ConvertMysqlTimeUnixTime(nullable.String)
		}

		if deadline.Valid {
			test.Deadline = conversion.ConvertMysqlTimeUnixTime(deadline.String)
		}

		if duration.Valid {
			test.Duration = int(duration.Int32)
		}

		test.DateCreated = conversion.ConvertMysqlTimeUnixTime(created)
		test.DateAssigned = conversion.ConvertMysqlTimeUnixTime(assigned)

		finalresult = append(finalresult, test)
	}

	return finalresult, err
}
