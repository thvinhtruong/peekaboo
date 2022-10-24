package mysqldb

import (
	"context"
	"server/app/domain/entity"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateTestClass(t *testing.T) {
	q := NewQuery("TestCreateTestClass")

	ctx := context.Background()
	for i := 1; i <= 10; i++ {
		err := q.AssignTestClass(ctx, entity.TestClassRelation{
			TestID:  i,
			ClassID: i,
		})
		require.NoError(t, err)
	}
}

func TestQueryTestClass(t *testing.T) {
	q := NewQuery("TestQueryTestClass")

	ctx := context.Background()
	for i := 6; i <= 12; i++ {
		res, err := q.QueryTestClass(ctx, i)
		require.NoError(t, err)
		require.NotEmpty(t, res)
	}
}

func TestDeleteTestClass(t *testing.T) {
	q := NewQuery("TestDeleteTestClass")

	ctx := context.Background()

	for i := 6; i <= 6; i++ {
		err := q.DeleteTestClass(ctx, i, i)
		require.NoError(t, err)
	}
}
