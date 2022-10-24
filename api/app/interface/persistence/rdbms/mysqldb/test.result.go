package mysqldb

import (
	"context"
	"database/sql"
	"errors"
	"server/app/domain/entity"
	"time"

	"github.com/jinzhu/copier"
)

func (q *Querier) CreateTestResult(ctx context.Context, data entity.TestResult) (int, error) {
	return insertTestResult(q, ctx, data)
}

func (q *Querier) UpdateTestResult(ctx context.Context, data entity.TestResult) error {
	return updateTestResult(q, ctx, data)
}

// Query TestResultIndex allows you to search according to flag, returns all fields.
func (q *Querier) QueryTestResultDetails(ctx context.Context, ID int) ([]entity.TestResult, error) {
	rows, err := q.DB.QueryContext(ctx, "SELECT * FROM testresults WHERE id = ? and active = ?", ID, 1)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return []entity.TestResult{}, err
	}

	data, err := refactorTestResultDetail(rows)
	if err != nil {
		return []entity.TestResult{}, err
	}

	if len(data) == 0 {
		return nil, nil
	}

	var result []entity.TestResult
	if err := copier.Copy(&result, &data[0]); err != nil {
		return []entity.TestResult{}, err
	}

	return result, nil
}

// Query TestResultIndex allows you to search according to flag, but only return the indexed value: ID, TestID, user_id, ClassID. Other fields are expected to be 0/nil.
// Flag determine the query element of the functions: [1 - TestClassID] [2 - user_id] [3 - DateCreated] [4 - TestClassID, user_id] [5 - TestClassID, DateCreated] [6 - user_id, DateCreated] [7 - TestClassID, user_id, DateCreated]
func (q *Querier) QueryTestResultIndexScore(ctx context.Context, TestClassID int, UserID int, DateCreate time.Time, Flag int) ([]entity.TestResult, error) {
	var err error
	var rows *sql.Rows
	switch Flag {
	case 1:
		rows, err = q.DB.QueryContext(ctx, "SELECT id, test_class_id, user_id, entity_code, datecreated, score, comment, resultnote FROM testresults WHERE test_class_id = ? and active = ?", TestClassID, 1)
	case 2:
		rows, err = q.DB.QueryContext(ctx, "SELECT id, test_class_id, user_id, entity_code, datecreated, score, comment, resultnote FROM testresults WHERE user_id = ? and active = ?", UserID, 1)
	case 3:
		rows, err = q.DB.QueryContext(ctx, "SELECT id, test_class_id, user_id, entity_code, datecreated, score, comment, resultnote FROM testresults WHERE datecreated = ? and active = ?", DateCreate, 1)
	case 4:
		rows, err = q.DB.QueryContext(ctx, "SELECT id, test_class_id, user_id, entity_code, datecreated, score, comment, resultnote FROM testresults WHERE test_class_id = ? AND user_id = ? AND active = ?", TestClassID, UserID, 1)
	case 5:
		rows, err = q.DB.QueryContext(ctx, "SELECT id, test_class_id, user_id, entity_code, datecreated, score, comment, resultnote FROM testresults WHERE test_class_id = ? AND datecreated = ? AND active = ?", TestClassID, DateCreate, 1)
	case 6:
		rows, err = q.DB.QueryContext(ctx, "SELECT id, test_class_id, user_id, entity_code, datecreated, score, comment, resultnote FROM testresults WHERE user_id = ? AND datecreated = ? AND active = ?", UserID, DateCreate, 1)
	case 7:
		rows, err = q.DB.QueryContext(ctx, "SELECT id, test_class_id, user_id, entity_code, datecreated, score, comment, resultnote FROM testresults WHERE test_class_id = ? AND user_id = ? AND datecreated = ? AND active = ?", TestClassID, UserID, DateCreate, 1)
	default:
		return nil, errors.New("error flag not found")
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	data, err := refactorTestResultIndex(rows)
	if err != nil {
		return nil, err
	}

	var result []entity.TestResult
	if err := copier.Copy(&result, data); err != nil {
		return nil, err
	}

	return result, nil
}

//DeleteTestResult can be use to delete a workdetail record of a student.
func (q *Querier) DeleteTestResult(ctx context.Context, ID int) error {
	_, err := q.DB.ExecContext(ctx, "delete from testresults where id = ?", ID)
	if err != nil {
		return err
	}

	return nil
}

func (q *Querier) ArchieveTestResult(ctx context.Context, TestClassID int, UserID int) error {
	_, err := q.DB.ExecContext(ctx, "UPDATE testresults set active = ? WHERE test_class_id = ? AND user_id = ?", TestClassID, UserID)
	if err != nil {
		return err
	}

	return nil
}

func (q *Querier) DeleteTestResultOfClass(ctx context.Context, TestClassID int) error {
	_, err := q.DB.ExecContext(ctx, "DELETE from testresults WHERE test_class_id = ?", TestClassID)
	if err != nil {
		return err
	}

	return nil
}
