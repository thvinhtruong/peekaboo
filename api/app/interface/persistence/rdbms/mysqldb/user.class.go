package mysqldb

import (
	"context"
	"database/sql"
	"strings"
)

// AddStudentToClass assigned a student to a class by create a pair of cid - sid in the db.
// Change to ClassID and UserID [].
func (q *Querier) AddUserClass(ctx context.Context, ClassID int, UserID int) error {
	stmt, err := q.DB.PrepareContext(ctx, "INSERT user_class SET uid = ?, cid = ?, active = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(UserID, ClassID, 1)
	if err != nil {
		if (strings.Index(err.Error(), "1062")) != 0 {
			return nil
		}

		return err
	}

	return nil
}

// QueryStudentInClass returns all the student id (sids) in the given class.
// Ref from QueryStudentInClass(cid int) ([]int, error)
// function return nil, nil if no student is currently in class.
func (q *Querier) QueryUserOfClass(ctx context.Context, ClassID int) ([]int, error) {
	rows, err := q.DB.QueryContext(ctx, "SELECT uid FROM user_class WHERE cid = ? and active = ?", ClassID, 1)

	if err == sql.ErrNoRows {
		return []int{}, nil
	}

	if err != nil {
		return nil, err
	}

	var result []int

	for rows.Next() {
		var tSID int
		err := rows.Scan(&tSID)
		if err != nil {
			return nil, err
		}

		result = append(result, tSID)
	}

	return result, nil
}

func (q *Querier) QueryClassOfUser(ctx context.Context, UserID int) ([]int, error) {
	rows, err := q.DB.QueryContext(ctx, "SELECT cid FROM user_class WHERE uid = ? and active = ?", UserID, 1)

	if err == sql.ErrNoRows {
		return []int{}, nil
	}

	if err != nil {
		return nil, err
	}

	var result []int

	for rows.Next() {
		var CID int
		err := rows.Scan(&CID)
		if err != nil {
			return nil, err
		}

		result = append(result, CID)
	}

	return result, nil
}

func (q *Querier) CheckExistedUserClass(ctx context.Context, UserID int, ClassID int) (existed bool, err error) {
	active := -1
	err = q.DB.QueryRowContext(ctx, "SELECT active FROM user_class WHERE uid = ? and cid = ?", UserID, ClassID).Scan(&active)
	if err == sql.ErrNoRows {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	if active == 0 {
		return true, nil
	}

	return false, nil
}

func (q *Querier) UnarchieveUserClass(ctx context.Context, UserID int, ClassID int) (err error) {
	stmt, err := q.DB.PrepareContext(ctx, "UPDATE user_class SET active = ? WHERE uid = ? AND cid = ?")
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, 1, UserID, ClassID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteStudentClass delete a sid-cid pair in the db, which means the student is not in the class from now.
func (q *Querier) DeleteUserClass(ctx context.Context, ClassID int, StudentID int) error {
	stmt, err := q.DB.PrepareContext(ctx, "UPDATE user_class SET active = ? WHERE uid = ? AND cid = ?")
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, 0, StudentID, ClassID)
	if err != nil {
		return err
	}

	return nil
}
