package mysqldb

import (
	"context"
	"database/sql"
	"encoding/json"
	"server/app/domain/entity"
	"server/utils/conversion"
	"time"
)

func (q *Querier) CreateSkillTest(ctx context.Context, st entity.SkillTest) (id int, err error) {
	stmt, err := q.DB.PrepareContext(ctx, "INSERT skill_tests SET media_url = ?, title = ?, content = ?, description = ?, type = ?, sections = ?, datecreated = ?")
	if err != nil {
		return 0, err
	}

	date := conversion.ConvertUnixTimeMySqlTime(time.Now().Unix())
	byteRes, err := json.Marshal(st.Content)
	if err != nil {
		return 0, err
	}

	result, err := stmt.ExecContext(ctx, st.MediaURL, st.Title, st.Content, st.Description, st.Type, (byteRes), date)
	if err != nil {
		return 0, err
	}

	k, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	id = int(k)

	return id, nil
}

func (q *Querier) QuerySkillTest(ctx context.Context, id int) (st entity.SkillTest, err error) {
	var created, jsonSection string
	var update sql.NullString
	err = q.DB.QueryRowContext(ctx, "SELECT * FROM skilltests WHERE id = ?", id).
		Scan(&st.Id, &st.MediaURL, &st.Title, &st.Content, &st.Description, &jsonSection, &st.Type, &created, &update)
	if err != nil {
		return st, err
	}

	if err = json.Unmarshal([]byte(jsonSection), &st.Section); err != nil {
		// panic(err)
		return st, err
	}

	return st, nil
}

func (q *Querier) UpdateSkillTest(ctx context.Context, st entity.SkillTest) (err error) {
	stmt, err := q.DB.PrepareContext(ctx, "INSERT skill_tests SET media_url = ?, title = ?, content = ?, description = ?, type = ?, sections = ?, dateupdated = ?")
	if err != nil {
		return err
	}

	date := conversion.ConvertUnixTimeMySqlTime(time.Now().Unix())
	byteRes, err := json.Marshal(st.Content)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, st.MediaURL, st.Title, st.Content, st.Description, st.Type, (byteRes), date)
	if err != nil {
		return err
	}

	return nil
}

func (q *Querier) DeleteSkillTest(ctx context.Context, st entity.SkillTest) (err error) {
	_, err = q.DB.ExecContext(ctx, "DELETE FROM skill_tests WHERE id = ?", 1)
	if err != nil {
		return err
	}

	return nil
}
