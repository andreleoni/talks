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
type ProductRepositoryImpl struct {
	pg database.DatabaseIface // Pode ser a conexão com postgres
}

// NewProductRepository initializes a new database structure
func NewProductRepository(db database.DatabaseIface) ProductRepositoryImpl {
	return ProductRepositoryImpl{db}
}

// Create simulates a create to the database
func (di ProductRepositoryImpl) Create(data string) error {
	return di.pg.Add("product:" + data)
}

// Delete simulates a delete to the database
func (di ProductRepositoryImpl) Delete(data string) error {
	return di.pg.Remove("product:" + data)
}
