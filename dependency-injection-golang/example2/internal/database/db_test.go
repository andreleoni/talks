package database

import (
	"example/mocks"
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func TestDB_Add(t *testing.T) {
	t.Run("when not returns errors", func(t *testing.T) {
		t.Helper()

		dbconnMock := mocks.PGConnecter{}
		db := NewDatabase(&dbconnMock)
		query := "my query"

		dbconnMock.On("Exec", query).Once().Return(nil)

		assert.Equal(t, db.Add(query), nil)
	})

	t.Run("when returns errors", func(t *testing.T) {
		t.Helper()

		dbconnMock := mocks.PGConnecter{}
		db := NewDatabase(&dbconnMock)
		query := "my query"

		expectedError := fmt.Errorf("anyone error")

		dbconnMock.On("Exec", query).Once().Return(expectedError)

		assert.Equal(t, db.Add(query), expectedError)
	})
}

func TestDB_Remove(t *testing.T) {
	t.Run("when not returns errors", func(t *testing.T) {
		t.Helper()

		dbconnMock := mocks.PGConnecter{}
		db := NewDatabase(&dbconnMock)
		query := "my query"

		dbconnMock.On("Exec", query).Once().Return(nil)

		assert.Equal(t, db.Remove(query), nil)
	})

	t.Run("when returns errors", func(t *testing.T) {
		t.Helper()

		dbconnMock := mocks.PGConnecter{}
		db := NewDatabase(&dbconnMock)
		query := "my query"

		expectedError := fmt.Errorf("anyone error")

		dbconnMock.On("Exec", query).Once().Return(expectedError)

		assert.Equal(t, db.Remove(query), expectedError)
	})
}
