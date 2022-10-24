package mysqldb

import (
	"context"
	"database/sql"
	"errors"
	"server/app/domain/entity"
	"server/utils/conversion"
	"time"
)

func (q *Querier) CreateTag(ctx context.Context, testtag entity.Tag) (int, error) {
	date := conversion.ConvertUnixTimeMySqlTime(time.Now().Unix())

	result, err := q.DB.ExecContext(ctx, "insert tags set tag = ?, info = ?, active = ?, datecreated = ?",
		testtag.Tag, testtag.Info, 1, date)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Replaceable fields: tag, info
func (q *Querier) UpdateTag(ctx context.Context, testtag entity.Tag) error {
	date := conversion.ConvertUnixTimeMySqlTime(time.Now().Unix())
	_, err := q.DB.ExecContext(ctx, "update tags set tag = ?, info = ?, dateupdated = ? where id = ?", testtag.Tag, testtag.Info, date, testtag.ID)
	if err != nil {
		return err
	}

	return nil
}

// Flag determine the query element of the functions: [0 - ID]  [1 - All]
func (q *Querier) QueryTag(ctx context.Context, TestTagID int, flag int) ([]entity.Tag, error) {
	switch flag {
	case 0:
		rows, err := q.DB.QueryContext(ctx, "select * from tags where id = ? and active = ?", TestTagID, 1)
		return refactorQueryTag(rows, err)
	case 1:
		rows, err := q.DB.QueryContext(ctx, "select * from tags where active = ?", 1)
		return refactorQueryTag(rows, err)
	default:
		return []entity.Tag{}, errors.New("error query tag")

	}
}

func refactorQueryTag(rows *sql.Rows, nerr error) (result []entity.Tag, err error) {
	if nerr == sql.ErrNoRows {
		return result, nil
	}

	if nerr != nil {
		return result, nerr
	}

	for rows.Next() {
		var nullable sql.NullString
		var tt entity.Tag
		var created string
		err = rows.Scan(&tt.ID, &tt.Tag, &tt.Info, &tt.Active, &created, &nullable)

		if err != nil {
			return result, err
		}

		if nullable.Valid {
			tt.DateUpdated = conversion.ConvertMysqlTimeUnixTime(nullable.String)
		}

		tt.DateCreated = conversion.ConvertMysqlTimeUnixTime(created)
		result = append(result, tt)
	}

	return result, nil
}

func (q *Querier) DeleteTag(ctx context.Context, TestTagID int) error {
	_, err := q.DB.ExecContext(ctx, "update tags set active = ? where id = ?", 0, TestTagID)
	if err != nil {
		return err
	}

	return nil
}
