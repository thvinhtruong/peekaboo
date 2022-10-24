package mysqldb

import (
	"context"
	"database/sql"
	"errors"
	"server/app/domain/entity"
	"server/utils/conversion"
	"time"
)

var (
	sql_insert = "INSERT users SET full_name = ?, username = ?, password = ?, gender = ?, address = ?, mail = ?, phone = ?, dob = ?, qualification = ?, entity_code = ?, active = ?, datecreated = ?"

	sql_update = "UPDATE users SET full_name = ?, username = ?, password = ?, gender = ?, address = ?, mail = ?, phone = ?, dob = ?, qualification = ?, dateupdated = ? WHERE active = ? and id = ?"

	sql_query_username = "SELECT * FROM users WHERE username = ? AND active = ?"

	sql_query_fullname = "SELECT * FROM users WHERE active = ? AND full_name like ?"

	sql_query_id = "SELECT * FROM users WHERE id = ? AND active = ?"
)

func insertUser(q *Querier, ctx context.Context, user entity.User) (int, error) {
	date := conversion.ConvertUnixTimeMySqlTime(time.Now().Unix())
	birthday := ""
	if user.Dob != 0 {
		birthday = conversion.ConvertUnixTimeMySqlTime(user.Dob)
	} else {
		birthday = conversion.ConvertUnixTimeMySqlTime(time.Now().Unix())
	}

	results, err := q.DB.ExecContext(ctx, sql_insert, user.FullName, user.Username, user.Password, user.Gender, user.Address, user.Mail, user.Phone, birthday, user.Qualification, user.EntityCode, 1, date)
	if err != nil {
		return 0, err
	}

	ID, err := results.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(ID), nil
}

func updateUser(q *Querier, ctx context.Context, user entity.User) error {
	date := conversion.ConvertUnixTimeMySqlTime(time.Now().Unix())
	bob := conversion.ConvertUnixTimeMySqlTime(user.Dob)

	_, err := q.DB.ExecContext(ctx, sql_update, user.FullName, user.Username, user.Password, user.Gender, user.Address, user.Mail, user.Phone, bob, user.Qualification, date, 1, user.ID)
	if err != nil {
		return err
	}

	return nil
}

// Do not select the Active, DateCreated and DateUpdated.
func refactorUserSelect(rows *sql.Rows, err error, HasPassword bool) (user []entity.User, nerr error) {
	if err == sql.ErrNoRows {
		return user, nerr
	}

	if err != nil {
		return []entity.User{}, err
	}

	for rows.Next() {
		var u entity.User
		var Address, Dob, Mail, Phone, Updated, Qualification sql.NullString
		var created string

		if err := rows.Scan(&u.ID, &u.FullName, &u.Username, &u.Password, &u.Gender, &Address,
			&Mail, &Phone, &Dob, &Qualification, &u.EntityCode, &u.Active, &created, &Updated); err != nil {
			return []entity.User{}, err
		}

		if Address.Valid {
			u.Address = Address.String
		}

		if Mail.Valid {
			u.Mail = Mail.String
		}

		if Phone.Valid {
			u.Phone = Phone.String
		}

		if Qualification.Valid {
			u.Qualification = Qualification.String
		}

		if !HasPassword {
			u.Password = ""
		}

		u.DateCreated = conversion.ConvertMysqlTimeUnixTime(created)

		if Updated.Valid {
			u.DateUpdated = conversion.ConvertMysqlTimeUnixTime(Updated.String)
		}

		if Dob.Valid {
			u.Dob = conversion.ConvertMysqlTimeUnixTime(Dob.String)
		}

		user = append(user, u)
	}

	return user, nil
}

func queryUser(q *Querier, ctx context.Context, Username string, Fullname string, ID int, Flag int, HasPassword bool) ([]entity.User, error) {
	switch Flag {
	case 0:
		if Username == "" {
			return []entity.User{}, errors.New("username must not empty")
		}

		rows, err := q.DB.QueryContext(ctx, sql_query_username, Username, 1)
		return refactorUserSelect(rows, err, HasPassword)

	case 1:
		if Fullname == "" {
			return []entity.User{}, errors.New("fullname must not empty")
		}

		rows, err := q.DB.QueryContext(ctx, sql_query_fullname, 1, "%"+Fullname+"%")
		return refactorUserSelect(rows, err, HasPassword)

	case 2:
		if ID == 0 {
			return []entity.User{}, errors.New("id must not empty")
		}

		rows, err := q.DB.QueryContext(ctx, sql_query_id, ID, 1)
		return refactorUserSelect(rows, err, HasPassword)

	case 3:
		rows, err := q.DB.QueryContext(ctx, "select * from users where active = ?", 1)
		return refactorUserSelect(rows, err, HasPassword)

	default:
		return []entity.User{}, errors.New("wrong Flag")
	}
}
