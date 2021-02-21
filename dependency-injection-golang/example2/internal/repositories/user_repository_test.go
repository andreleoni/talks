package repositories

import (
	"example/mocks"
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func TestUserRepository_Create(t *testing.T) {
	t.Run("when not returns errors", func(t *testing.T) {
		t.Helper()

		dbMock := mocks.DatabaseIface{}
		userRepository := NewUserRepository(&dbMock)
		query := "my query"

		dbMock.On("Add", "user:"+query).Once().Return(nil)

		assert.Equal(t, userRepository.Create(query), nil)
	})

	t.Run("when returns errors", func(t *testing.T) {
		t.Helper()

		dbMock := mocks.DatabaseIface{}
		userRepository := NewUserRepository(&dbMock)
		query := "my query"

		expectedError := fmt.Errorf("anyone error")

		dbMock.On("Add", "user:"+query).Once().Return(expectedError)

		assert.Equal(t, userRepository.Create(query), expectedError)
	})
}

func TestUserRepository_Delete(t *testing.T) {
	t.Run("when not returns errors", func(t *testing.T) {
		t.Helper()

		dbMock := mocks.DatabaseIface{}
		userRepository := NewUserRepository(&dbMock)
		query := "my query"

		dbMock.On("Remove", "user:"+query).Once().Return(nil)

		assert.Equal(t, userRepository.Delete(query), nil)
	})

	t.Run("when returns errors", func(t *testing.T) {
		t.Helper()

		dbMock := mocks.DatabaseIface{}
		userRepository := NewUserRepository(&dbMock)
		query := "my query"

		expectedError := fmt.Errorf("anyone error")

		dbMock.On("Remove", "user:"+query).Once().Return(expectedError)

		assert.Equal(t, userRepository.Delete(query), expectedError)
	})
}
