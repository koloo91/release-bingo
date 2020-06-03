package main

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/koloo91/release-bingo/controller"
	"github.com/koloo91/release-bingo/repository"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

const (
	host     = "localhost" // TODO: make it configurable
	port     = 5432
	user     = "postgres"
	password = "Pass00"
	dbname   = "postgres"
)

func main() {

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Error openening sql connection: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Error creating migration driver instance: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://./migrations", "postgres", driver)
	if err != nil {
		log.Fatalf("Error creating migration database instance: %v", err)
	}

	if err := m.Up(); err != nil {
		log.Printf("Error running migrations: %v", err)
	}

	repository.SetDatabase(db)

	router := controller.SetupRoutes()

	server := http.Server{
		Addr:         ":9000", // TODO: make it configurable
		Handler:      router,
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}

	log.Println("Starting server at port :9000")
	log.Fatal(server.ListenAndServe())
}
