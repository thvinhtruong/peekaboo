package mysqldb

import (
	"context"
	"database/sql"
	"server/app/domain/entity"
	"server/utils/conversion"
	"time"
)

// Returns the ID of the class.
func (q *Querier) CreateClass(ctx context.Context, c entity.Class) (int, error) {
	stmt, err := q.DB.PrepareContext(ctx, "INSERT classes SET class_name = ?, info = ?, announcement = ?, room_code = ?, level = ?, active = ? , datecreated = ?")
	if err != nil {
		return 0, err
	}

	date := conversion.ConvertUnixTimeMySqlTime(time.Now().Unix())
	result, err := stmt.ExecContext(ctx, c.Classname, c.Info, c.Announcement, c.RoomCode, c.Level, 1, date)
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

// Query prior to ClassID, then ClassName
func (q *Querier) QueryClass(ctx context.Context, ClassID int, ClassName string) ([]entity.Class, error) {
	if ClassID != 0 {
		rows, err := q.DB.QueryContext(ctx, "SELECT * FROM classes WHERE id = ? and active = ?", ClassID, 1)
		return refactorQueryClass(rows, err)
	}

	if len(ClassName) != 0 {
		rows, err := q.DB.QueryContext(ctx, "SELECT * FROM classes WHERE class_name like ? and active = ?", "%"+ClassName+"%", 1)
		return refactorQueryClass(rows, err)
	}

	return []entity.Class{}, nil
}

// Query the last X classes in the database.
func (q *Querier) QueryAllClass(ctx context.Context) ([]entity.Class, error) {
	rows, err := q.DB.QueryContext(ctx, "SELECT * FROM classes ORDER BY id DESC")
	return refactorQueryClass(rows, err)
}

func refactorQueryClass(rows *sql.Rows, nerr error) (result []entity.Class, err error) {
	if nerr == sql.ErrNoRows {
		return result, nil
	}

	if nerr != nil {
		return result, nerr
	}

	for rows.Next() {
		var class entity.Class
		var nullable sql.NullString
		var created string

		err = rows.Scan(&class.ID, &class.Classname, &class.Info, &class.Announcement, &class.RoomCode, &class.Level, &class.Active, &created, &nullable)
		if err == sql.ErrNoRows {
			return []entity.Class{}, nil
		}

		if err != nil {
			return []entity.Class{}, err
		}

		class.DateCreated = conversion.ConvertMysqlTimeUnixTime(created)

		if nullable.Valid {
			class.DateUpdated = conversion.ConvertMysqlTimeUnixTime(nullable.String)
		}

		result = append(result, class)
	}

	return result, nil
}

// UpdateClass allow teacher to update the class information.
func (q *Querier) UpdateClass(ctx context.Context, c entity.Class) error {
	stmt, err := q.DB.PrepareContext(ctx, "UPDATE classes SET class_name = ?, info = ?, announcement = ?, room_code = ?, level = ?, dateupdated = ? WHERE id = ? and active = ?")
	if err != nil {
		return err
	}

	date := conversion.ConvertUnixTimeMySqlTime(time.Now().Unix())

	_, err = stmt.ExecContext(ctx, c.Classname, c.Info, c.Announcement, c.RoomCode, c.Level, date, c.ID, 1)
	if err != nil {
		return err
	}

	return nil
}

// ArchiveClass changes the active status to 0.
func (q *Querier) ArchiveClass(ctx context.Context, ID int) error {
	date := conversion.ConvertUnixTimeMySqlTime(time.Now().Unix())
	_, err := q.DB.ExecContext(ctx, "UPDATE classes set active = ?, dateupdated = ? where id = ?", 0, date, ID)
	if err != nil {
		return err
	}

	return nil
}

// ArchiveClass changes the active status to 0.
func (q *Querier) DeleteClass(ctx context.Context, ID int) error {
	_, err := q.DB.ExecContext(ctx, "UPDATE classes SET active = ? where id = ?", 0, ID)
	if err != nil {
		return err
	}

	return nil
}
