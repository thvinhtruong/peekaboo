package mysqldb

import (
	"context"
	"log"
	"server/app/domain/entity"
	"server/utils/service"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateSkillTest(t *testing.T) {
	q := NewQuery("TestCreateSkillTest")
	ctx := context.Background()
	for i := 1; i <= 10; i++ {
		b := entity.SkillTest{
			MediaURL:    service.RandomString(10),
			Title:       service.RandomString(10),
			Content:     service.RandomString(10),
			Description: service.RandomString(10),
		}
		id, err := q.CreateSkillTest(ctx, b)
		if err != nil {
			log.Fatalf("fail with error: %v", err)
		}
		require.NoError(t, err)
		require.NotEqual(t, 0, id)
	}
}

func TestQuerySkillTest(t *testing.T) {
	q := NewQuery("TestQuerySkillTest")
	ctx := context.Background()
	res, err := q.QuerySkillTest(ctx, 1)
	require.NoError(t, err)
	require.NotEmpty(t, res)
}

func TestUpdateSkillTest(t *testing.T) {
	q := NewQuery("TestUpdateSkillTest")
	ctx := context.Background()
	res, err := q.QuerySkillTest(ctx, 1)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	err = q.UpdateSkillTest(ctx, res)
	require.NoError(t, err)

}

func TestDeleteSkillTest(t *testing.T) {
	q := NewQuery("TestDeleteSkillTest")
	ctx := context.Background()
	res, err := q.QuerySkillTest(ctx, 1)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	err = q.DeleteSkillTest(ctx, res)
	require.NoError(t, err)
}
