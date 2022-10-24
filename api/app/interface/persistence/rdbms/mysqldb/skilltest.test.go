package mysqldb

import (
	"context"
	"database/sql"
)

func (q *Querier) AssignSkillTest2Test(ctx context.Context, TestID int, SkillTestID int) error {
	stmt, err := q.DB.PrepareContext(ctx, "INSERT skilltest_test SET tid = ?, stid = ?")
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, TestID, SkillTestID)
	return err
}

func (q *Querier) QuerySkillTestOfTest(ctx context.Context, TestID int) (result []int, err error) {
	rows, err := q.DB.QueryContext(ctx, "SELECT stid FROM skilltest_test WHERE tid = ? ", TestID)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tid int
		err := rows.Scan(&tid)
		if err != nil {
			return result, err
		}

		result = append(result, tid)
	}

	return result, nil
}

func (q *Querier) DeleteSkillTestAndTest(ctx context.Context, TestID int, SkillTestID int) error {
	_, err := q.DB.ExecContext(ctx, "DELETE FROM skilltest_test WHERE tid = ? AND stid = ?", TestID, SkillTestID)
	if err != nil {
		return err
	}

	return nil
}
