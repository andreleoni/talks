package main

import (
	"example/internal/application"
	"example/internal/database"
	"example/internal/repositories"
	"example/pkg/pgconn"
)

func main() {
	// Simula buscar a conexão de um serviço externo
	pgConn := pgconn.PGConn{}

	// Inicializa dependencias
	currentDB := database.NewDatabase(pgConn)

	// Use the same database to each repository
	userRepository := repositories.NewUserRepository(currentDB)
	productRepository := repositories.NewProductRepository(currentDB)

	application := application.NewApplication(userRepository, productRepository)

	// Process the application
	application.Process()
}
