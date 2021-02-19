package repositories

import (
	"example/internal/database"
)

// UsersRepositoryIface is the interface for users repository structure
type UsersRepositoryIface interface {
	Create(string) error
	Delete(string) error
}

// usersRepositoryImpl is the impl to user repository
type usersRepositoryImpl struct {
	pg database.DatabaseIface // Pode ser a conexão com postgres
}

// NewUserRepository initializes a new database structure
func NewUserRepository(db database.DatabaseIface) UsersRepositoryIface {
	return usersRepositoryImpl{db}
}

// Create simulates a create to the database
func (di usersRepositoryImpl) Create(data string) error {
	return di.pg.Add("user:" + data)
}

// Delete simulates a delete to the database
func (di usersRepositoryImpl) Delete(data string) error {
	return di.pg.Remove("user:" + data)
}
