package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"go-grpc/internal/pb"
	"go-grpc/internal/repository"
	"go-grpc/internal/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	database, err := sql.Open("sqlite3", "file::memory:?mode=memory&cache=shared")
	if err != nil {
		panic(err)
	}
	defer database.Close()

	_, err = database.Exec(`
		CREATE TABLE IF NOT EXISTS categories
		(
			id          VARCHAR(36)  NOT NULL PRIMARY KEY,
			name        VARCHAR(255) NOT NULL,
			description MEDIUMTEXT   NOT NULL,
			created_at  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		panic(err)
	}

	categoryRepository := repository.NewSqlCategoryRepository(database)
	categoryService := rpc.NewCategoryService(categoryRepository)

	server := grpc.NewServer()
	reflection.Register(server)
	pb.RegisterCategoryServiceServer(server, categoryService)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	err = server.Serve(listener)
	if err != nil {
		panic(err)
	}
}
