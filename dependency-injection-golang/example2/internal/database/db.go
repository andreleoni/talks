package database

import "example/pkg/pgconn"

// DatabaseIface is the interface for database structure
type DatabaseIface interface {
	Add(string) error
	Remove(string) error
}

// DatabaseImpl is the impl to database structure
type DatabaseImpl struct {
	pg pgconn.PGConnecter // Pode ser a conexão com postgres
}

// NewDatabase initializes a new database structure
func NewDatabase(db pgconn.PGConnecter) DatabaseImpl {
	return DatabaseImpl{db}
}

// Add simulates a exec of database
func (di DatabaseImpl) Add(data string) error {
	return di.pg.Exec(data)
}

// Remove simulates a exec of database
func (di DatabaseImpl) Remove(data string) error {
	return di.pg.Exec(data)
}
