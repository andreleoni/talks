package pgconn

import "fmt"

// This package only simluate a external PG connection library

// PGConnecter is the interface for pg conn
type PGConnecter interface {
	Exec(query string) error
}

// PGConn Simula um pacote externo
type PGConn struct{}

// Exec simula a execução de uma query
func (pg PGConn) Exec(query string) error {
	fmt.Println(query)

	return nil
}
