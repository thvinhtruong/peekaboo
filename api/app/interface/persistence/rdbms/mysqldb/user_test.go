package mysqldb

import (
	"context"
	"log"
	"server/app/domain/entity"
	"server/utils/service"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	q := NewQuery("TestCreateUser")
	ctx := context.Background()
	for i := 1; i <= 10; i++ {
		a := entity.User{
			FullName:      "NicholasP",
			Username:      service.RandomString(5),
			Password:      "ducthangng",
			Gender:        "Male",
			Address:       "8/8B 38 str, 8, TDCity",
			Dob:           time.Now().Unix(),
			Phone:         service.RandomString(5),
			Mail:          "something you should not know",
			Qualification: "proplayer",
			EntityCode:    1,
		}

		id, err := q.CreateUser(ctx, a)
		if err != nil {
			log.Fatalf("fail with error: %v", err)
		}
		require.NoError(t, err)
		require.NotEqual(t, 0, id)

		// Student
		a.EntityCode = 2
		a.Username = service.RandomString(5)
		a.Phone = service.RandomString(5)
		id, err = q.CreateUser(ctx, a)
		if err != nil {
			log.Fatalf("fail with error: %v", err)
		}
		require.NoError(t, err)
		require.NotEqual(t, 0, id)

		// Teacher
		a.EntityCode = 3
		a.Username = service.RandomString(5)
		a.Phone = service.RandomString(5)
		id, err = q.CreateUser(ctx, a)
		if err != nil {
			log.Fatalf("fail with error: %v", err)
		}
		require.NoError(t, err)
		require.NotEqual(t, 0, id)
	}
}

func TestUpdateUser(t *testing.T) {
	q := NewQuery("TestUpdateUser")

	ctx := context.Background()

	for i := 1; i <= 30; i++ {
		res, err := q.QueryUser(ctx, "", "", i, 2, false)
		require.NoError(t, err)
		require.NotEmpty(t, res)
		require.Equal(t, res[0].ID, i)
		require.Equal(t, res[0].Password, "")

		res, err = q.QueryUser(ctx, "", "", i, 2, true)
		require.NoError(t, err)
		require.NotEmpty(t, res)
		require.Equal(t, res[0].ID, i)
		require.NotEqual(t, res[0].Password, "")

		res[0].Username = res[0].Username + " 1"
		err = q.UpdateUser(ctx, res[0])
		require.NoError(t, err)

		res1, err := q.QueryUser(ctx, "", res[0].FullName, i, 1, true)
		require.NoError(t, err)
		require.NotEmpty(t, res1)
		require.NotEqual(t, nil, res1)
		require.NotNil(t, res1[0].Password)

		res2, err := q.QueryUser(ctx, res[0].Username, "", i, 0, true)
		require.NoError(t, err)
		require.NotEmpty(t, res1)
		require.NotEqual(t, nil, res1)
		require.NotNil(t, res2[0].Password)
	}
}

func TestDeleteUser(t *testing.T) {
	q := NewQuery("TestDeleteUser")

	ctx := context.Background()
	for i := 1; i <= 4; i++ {
		err := q.DeleteUser(ctx, i)
		require.NoError(t, err)
	}
}
