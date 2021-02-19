package application

import (
	"example/internal/repositories"
)

// Application is the impl to user repository
type Application struct {
	userRepository    repositories.UsersRepositoryIface
	productRepostiory repositories.ProductRepositoryIface
}

// NewApplication initializes a new application structure
func NewApplication(
	userRepository repositories.UsersRepositoryIface,
	productRepostiory repositories.ProductRepositoryIface) Application {
	return Application{userRepository, productRepostiory}
}

// Process simulates a create and delete for user repository
func (a Application) Process() {
	a.userRepository.Create("Simulating a user creation...")

	a.userRepository.Delete("Simulating a user delete...")

	a.productRepostiory.Create("Simulating a product creation...")

	a.productRepostiory.Delete("Simulating a product delete...")
}
