package mysqldb

import (
	"context"
	"database/sql"
	"server/app/domain/entity"
)

func (q *Querier) AssignTestClass(ctx context.Context, TestClass entity.TestClassRelation) error {
	stmt, err := q.DB.PrepareContext(ctx, "INSERT test_class SET tid = ?, cid = ?")
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, TestClass.TestID, TestClass.ClassID)
	if err != nil {
		return err
	}

	return nil
}

func (q *Querier) QueryTestClass(ctx context.Context, TestClassID int) (result []entity.TestClassRelation, err error) {
	rows, err := q.DB.QueryContext(ctx, "SELECT cid, tid FROM test_class WHERE id = ? ", TestClassID)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tid, cid int
		err := rows.Scan(&cid, &tid)
		if err != nil {
			return result, err
		}

		result = append(result, entity.TestClassRelation{ID: TestClassID, TestID: tid, ClassID: cid})
	}

	return result, nil
}

func (q *Querier) QueryTestOfClass(ctx context.Context, ClassID int) (result []entity.TestClassRelation, err error) {
	rows, err := q.DB.QueryContext(ctx, "SELECT id, tid FROM test_class WHERE cid = ? ", ClassID)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id, tid int
		err := rows.Scan(&id, &tid)
		if err != nil {
			return result, err
		}

		result = append(result, entity.TestClassRelation{ID: id, TestID: tid, ClassID: ClassID})
	}

	return result, nil
}

func (q *Querier) QueryClassDoneTest(ctx context.Context, TestID int) (result []entity.TestClassRelation, err error) {
	rows, err := q.DB.QueryContext(ctx, "SELECT id, cid FROM test_class WHERE tid = ?", TestID)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id, cid int
		err := rows.Scan(&id, &cid)
		if err != nil {
			return result, err
		}

		result = append(result, entity.TestClassRelation{ID: id, TestID: TestID, ClassID: cid})
	}

	return result, nil
}

func (q *Querier) DeleteTestClass(ctx context.Context, TestID int, ClassID int) error {
	_, err := q.DB.ExecContext(ctx, "DELETE FROM test_class WHERE tid = ? AND cid = ?", TestID, ClassID)
	if err != nil {
		return err
	}

	return nil
}
