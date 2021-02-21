package repositories

import (
	"example/mocks"
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func TestProductRepository_Create(t *testing.T) {
	t.Run("when not returns errors", func(t *testing.T) {
		t.Helper()

		dbMock := mocks.DatabaseIface{}
		productRepository := NewProductRepository(&dbMock)
		query := "my query"

		dbMock.On("Add", "product:"+query).Once().Return(nil)

		assert.Equal(t, productRepository.Create(query), nil)
	})

	t.Run("when returns errors", func(t *testing.T) {
		t.Helper()

		dbMock := mocks.DatabaseIface{}
		productRepository := NewProductRepository(&dbMock)
		query := "my query"

		expectedError := fmt.Errorf("anyone error")

		dbMock.On("Add", "product:"+query).Once().Return(expectedError)

		assert.Equal(t, productRepository.Create(query), expectedError)
	})
}

func TestProductRepository_Delete(t *testing.T) {
	t.Run("when not returns errors", func(t *testing.T) {
		t.Helper()

		dbMock := mocks.DatabaseIface{}
		productRepository := NewProductRepository(&dbMock)
		query := "my query"

		dbMock.On("Remove", "product:"+query).Once().Return(nil)

		assert.Equal(t, productRepository.Delete(query), nil)
	})

	t.Run("when returns errors", func(t *testing.T) {
		t.Helper()

		dbMock := mocks.DatabaseIface{}
		productRepository := NewProductRepository(&dbMock)
		query := "my query"

		expectedError := fmt.Errorf("anyone error")

		dbMock.On("Remove", "product:"+query).Once().Return(expectedError)

		assert.Equal(t, productRepository.Delete(query), expectedError)
	})
}
