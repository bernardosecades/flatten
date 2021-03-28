package main

import (
	"fmt"
	"github.com/bernardosecades/flatten/internal/server/http"
	"github.com/bernardosecades/flatten/internal/storage/postgres"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

var commitHash string

func init() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Print("Not .env file found")
	}

	fmt.Printf("Build Time: %s\n", time.Now().Format(time.RFC3339))
	fmt.Printf("Version: %s\n", commitHash)
}

func main() {

	port := os.Getenv("SERVER_PORT")
	httpAddr := fmt.Sprintf(":%s", port)

	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")
	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	repository := postgres.NewPostgresRepository(dbHost, dbPort, dbUser, dbPass, dbName)
	handler := http.NewHandler(repository)

	server := http.NewServer(handler, httpAddr)
	err := server.Serve()

	if err != nil {
		log.Fatal(err)
	}
}
