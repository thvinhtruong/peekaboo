package mysqldb

import (
	"context"
	"server/app/domain/entity"
	"server/utils/service"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateInternalTest(t *testing.T) {
	q := NewQuery("TestCreateTest")

	ctx := context.Background()
	for i := 1; i <= 20; i++ {
		test := entity.Test{
			TagID:            11,
			TestName:         service.RandomString(5) + " " + service.RandomString(2),
			Info:             service.RandomString(5),
			DateAssigned:     time.Now().Unix(),
			Deadline:         time.Now().Unix(),
			TargetEntityCode: service.RandomNumber(1, 4),
			CreatedUserID:    service.RandomNumber(5, 30),
		}

		res, err := q.CreateTest(ctx, test)
		require.NoError(t, err)
		require.NotEqual(t, nil, res)
	}
}

func TestQueryTest(t *testing.T) {
	q := NewQuery("TestCreateTest")
	ctx := context.Background()

	for i := 3; i <= 30; i++ {
		res1, err := q.QueryTestHeadline(ctx, i, "")
		require.NoError(t, err)
		require.NotEmpty(t, res1)
		require.NotNil(t, res1)

		res2, err := q.QueryTestHeadline(ctx, 0, res1[0].TestName)
		require.NoError(t, err)
		require.NotEmpty(t, res2)
		require.NotNil(t, res2)

		require.Equal(t, res1[0], res2[0])
	}

}

func TestQueryAllTests(t *testing.T) {
	q := NewQuery("TestQueryTest")

	ctx := context.Background()
	_, err := q.QueryAllTestHeadlines(ctx)
	require.NoError(t, err)
}

func TestUpdateTest(t *testing.T) {
	q := NewQuery("TestUpdateTest")
	ctx := context.Background()

	for i := 3; i <= 30; i++ {
		res1, err := q.QueryTestHeadline(ctx, i, "")
		require.NoError(t, err)
		require.NotEmpty(t, res1)
		require.NotNil(t, res1)

		res1[0].Info = "gahahahaha"
		err = q.UpdateTest(ctx, res1[0])
		require.NoError(t, err)
	}
}

func TestDeletePTest(t *testing.T) {
	q := NewQuery("TestDeleteTest")
	ctx := context.Background()

	for i := 1; i <= 5; i++ {
		err := q.DeleteTest(ctx, i)
		require.NoError(t, err)
	}
}
