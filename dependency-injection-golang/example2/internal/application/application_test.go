package application

import (
	"example/mocks"
	"testing"
)

func TestApplication_Process(t *testing.T) {
	userRespositoryMock := mocks.UsersRepositoryIface{}
	productRespositoryMock := mocks.ProductRepositoryIface{}

	app := NewApplication(&userRespositoryMock, &productRespositoryMock)

	userRespositoryMock.On("Create", "Simulating a user creation...").Once().Return(nil)

	userRespositoryMock.On("Delete", "Simulating a user delete...").Once().Return(nil)

	productRespositoryMock.On("Create", "Simulating a product creation...").Once().Return(nil)

	productRespositoryMock.On("Delete", "Simulating a product delete...").Once().Return(nil)

	app.Process()
}
