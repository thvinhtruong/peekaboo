package mysqldb

import (
	"context"
	"server/app/domain/entity"
	"server/utils/conversion"
	"time"
)

func (q *Querier) CreateUser(ctx context.Context, user entity.User) (int, error) {
	return insertUser(q, ctx, user)
}

// Flag determine the query element of the functions: [0 - username]  [1 - fullname] [2 - ID] [3 - All].
// Option 2 is unindex search, therefore do not overuse.
func (q *Querier) QueryUser(ctx context.Context, Username string, Fullname string, ID int, Flag int, HasPassword bool) ([]entity.User, error) {
	return queryUser(q, ctx, Username, Fullname, ID, Flag, HasPassword)
}

func (q *Querier) UpdateUser(ctx context.Context, user entity.User) error {
	return updateUser(q, ctx, user)
}

func (q *Querier) DeleteUser(ctx context.Context, ID int) error {
	date := conversion.ConvertUnixTimeMySqlTime(time.Now().Unix())
	_, err := q.DB.ExecContext(ctx, "update users set active = ?, dateupdated = ? where id = ?", 0, date, ID)
	if err != nil {
		return err
	}

	return nil
}
