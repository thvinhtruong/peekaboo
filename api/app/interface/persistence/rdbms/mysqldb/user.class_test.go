package mysqldb

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddStudentClass(t *testing.T) {
	q := NewQuery("TestAddStudentToClass")

	ctx := context.Background()
	for i := 1; i <= 15; i++ {
		err := q.AddUserClass(ctx, i, i+1)
		require.NoError(t, err)
	}
}

func TestQueryStudentClass(t *testing.T) {
	q := NewQuery("TestQueryStudentInClass")

	ctx := context.Background()
	for i := 1; i <= 10; i++ {
		students, err := q.QueryUserOfClass(ctx, i)
		require.NoError(t, err)
		require.NotEmpty(t, students)
		require.Equal(t, students[0], i+1)
	}
}

func TestDeleteStudentClass(t *testing.T) {
	q := NewQuery("TestDeleteStudentClass")
	ctx := context.Background()

	for i := 1; i <= 10; i++ {
		result, err := q.QueryUserOfClass(ctx, i)
		require.NoError(t, err)

		for _, v := range result {
			err = q.DeleteUserClass(ctx, i, v)
			require.NoError(t, err)
		}
	}
}
