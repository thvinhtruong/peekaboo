package mysqldb

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAssignSkillTest2Test(t *testing.T) {
	q := NewQuery("TestAssignSkillTest2Test")
	ctx := context.Background()
	err := q.AssignSkillTest2Test(ctx, 1, 1)
	require.NoError(t, err)
}

func TestQuerySkillTestOfTest(t *testing.T) {
	q := NewQuery("TestQuerySkillTestOfTest")
	ctx := context.Background()
	res, err := q.QuerySkillTestOfTest(ctx, 1)
	require.NoError(t, err)
	require.NotEmpty(t, res)
}

func TestDeleteSkillTestAndTest(t *testing.T) {
	q := NewQuery("TestDeleteSkillTestAndTest")
	ctx := context.Background()
	err := q.DeleteSkillTestAndTest(ctx, 1, 1)
	require.NoError(t, err)
}
