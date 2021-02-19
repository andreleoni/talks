package repositories

import (
	"example/internal/database"
)

// ProductRepositoryIface is the interface for Product repository structure
type ProductRepositoryIface interface {
	Create(string) error
	Delete(string) error
}

// ProductRepositoryImpl is the impl to Product repository
type productRepositoryImpl struct {
	pg database.DatabaseIface // Pode ser a conexão com postgres
}

// NewProductRepository initializes a new database structure
func NewProductRepository(db database.DatabaseIface) ProductRepositoryIface {
	return productRepositoryImpl{db}
}

// Create simulates a create to the database
func (di productRepositoryImpl) Create(data string) error {
	return di.pg.Add("product:" + data)
}

// Delete simulates a delete to the database
func (di productRepositoryImpl) Delete(data string) error {
	return di.pg.Remove("product:" + data)
}
