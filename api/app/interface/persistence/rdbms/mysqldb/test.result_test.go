package mysqldb

import (
	"context"
	"log"
	"server/app/domain/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateTestResult(t *testing.T) {
	q := NewQuery("TestCreateTestResult")
	ctx := context.Background()
	for i := 1; i <= 10; i++ {
		body := entity.TestResult{
			TestClassID: 1,
			UserID:      1,
			Score:       8,
			DateCreated: 23,
			Comment:     "good",
		}
		id, err := q.CreateTestResult(ctx, body)
		if err != nil {
			log.Fatalf("fail with error: %v", err)
		}
		require.NoError(t, err)
		require.NotEqual(t, 0, id)
	}
}

func TestUpdateTestResult(t *testing.T) {
	q := NewQuery("TestUpdateTestResult")
	ctx := context.Background()

	body := entity.TestResult{
		TestClassID: 1,
		UserID:      1,
		Score:       8,
		DateCreated: 23,
		Comment:     "good",
	}

	err := q.UpdateTestResult(ctx, body)
	require.NoError(t, err)
}

func TestQueryTestResultDetails(t *testing.T) {
	q := NewQuery("TestQueryTestResultDetails")
	ctx := context.Background()
	res, err := q.QueryTestResultDetails(ctx, 1)
	require.NoError(t, err)
	require.NotEmpty(t, res)
}

func TestQueryTestResultIndexScore(t *testing.T) {
	q := NewQuery("TestQueryTestResultIndexScore")
	ctx := context.Background()

	for i := 1; i <= 7; i++ {
		res, err := q.QueryTestResultIndexScore(ctx, 1, 1, time.Now(), i)
		require.NoError(t, err)
		require.NotEmpty(t, res)
	}
}

func TestDeleteTestResult(t *testing.T) {
	q := NewQuery("TestDeleteTestResult")
	ctx := context.Background()
	res, err := q.QueryTestResultDetails(ctx, 1)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	for _, item := range res {
		err = q.DeleteTestResult(ctx, item.ID)
		require.NoError(t, err)
	}
}

func TestArchieveTestResult(t *testing.T) {
	q := NewQuery("TestArchiveTestResult")
	ctx := context.Background()
	res, err := q.QueryTestResultDetails(ctx, 1)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	for _, item := range res {
		err = q.ArchieveTestResult(ctx, item.TestClassID, item.UserID)
		require.NoError(t, err)
	}
}

func TestDeleteTestResultOfClass(t *testing.T) {
	q := NewQuery("TestDeleteResultOfClass")
	ctx := context.Background()
	res, err := q.QueryTestResultDetails(ctx, 1)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	for _, item := range res {
		err = q.DeleteTestResultOfClass(ctx, item.TestClassID)
		require.NoError(t, err)
	}
}
