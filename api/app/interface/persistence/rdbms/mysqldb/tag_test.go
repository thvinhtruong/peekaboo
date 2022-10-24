package mysqldb

import (
	"context"
	"log"
	"server/app/domain/entity"
	"testing"

	"server/utils/service"

	"github.com/stretchr/testify/require"
)

func TestCreateTestTag(t *testing.T) {
	q := NewQuery("TestCreateTestTag")
	ctx := context.Background()
	for i := 1; i <= 10; i++ {
		b := entity.Tag{
			Tag:    service.RandomString(10),
			Info:   service.RandomString(10),
			Active: 1,
		}

		id, err := q.CreateTag(ctx, b)
		if err != nil {
			log.Fatalf("fail with error: %v", err)
		}

		require.NoError(t, err)
		require.NotEqual(t, 0, id)
	}
}

func TestUpdateTestTag(t *testing.T) {
	q := NewQuery("TestUpdateTestTag")

	ctx := context.Background()

	for i := 1; i <= 10; i++ {
		res, err := q.QueryTag(ctx, i, 0)
		require.NoError(t, err)
		require.NotEmpty(t, res)

		res[0].Info = "Just been changed to a very fucking long description"
		err = q.UpdateTag(ctx, res[0])
		require.NoError(t, err)
	}
}

func TestQueryAllTestTag(t *testing.T) {
	q := NewQuery("TestQueryAllTestTag")

	ctx := context.Background()

	result, err := q.QueryTag(ctx, 0, 1)
	if err != nil {
		log.Fatalf("fail with error: %v", err)
	}

	require.NoError(t, err)
	require.NotNil(t, result, t)
}

func TestQueryTestTag(t *testing.T) {
	q := NewQuery("TestQueryTestTag")

	ctx := context.Background()
	tags, err := q.QueryTag(ctx, 0, 1)
	require.NoError(t, err)

	for _, tag := range tags {
		res1, err := q.QueryTag(ctx, tag.ID, 0)

		require.NoError(t, err)
		require.Equal(t, tag, res1[0])
	}
}

func TestDeleteTestTag(t *testing.T) {
	q := NewQuery("TestDeleteTestTag")

	ctx := context.Background()
	for i := 1; i <= 10; i++ {
		err := q.DeleteTag(ctx, i)
		require.NoError(t, err)
	}
}
