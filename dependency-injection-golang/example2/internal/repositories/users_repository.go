package repositories

import (
	"example/internal/database"
)

// UsersRepositoryIface is the interface for users repository structure
type UsersRepositoryIface interface {
	Create(string) error
	Delete(string) error
}

// UsersRepositoryImpl is the impl to user repository
type UsersRepositoryImpl struct {
	pg database.DatabaseIface // Pode ser a conexão com postgres
}

// NewUserRepository initializes a new database structure
func NewUserRepository(db database.DatabaseIface) UsersRepositoryImpl {
	return UsersRepositoryImpl{db}
}

// Create simulates a create to the database
func (di UsersRepositoryImpl) Create(data string) error {
	return di.pg.Add("user:" + data)
}

// Delete simulates a delete to the database
func (di UsersRepositoryImpl) Delete(data string) error {
	return di.pg.Remove("user:" + data)
}
