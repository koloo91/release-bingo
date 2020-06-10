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
	"os"
	"strings"
	"time"
)

var (
	serverPort = getEnvOrDefault("SERVER_PORT", "9000")

	host     = getEnvOrDefault("DB_HOST", "localhost")
	port     = getEnvOrDefault("DB_PORT", "5432")
	user     = getEnvOrDefault("DB_USER", "postgres")
	password = getEnvOrDefault("DB_PASSWORD", "Pass00")
	dbname   = getEnvOrDefault("DB_NAME", "postgres")

	users = getEnvOrDefault("USERS", "admin:admin")
)

func main() {

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
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

	db.SetConnMaxLifetime(1 * time.Hour)
	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(2)

	repository.SetDatabase(db)

	accounts := make(map[string]string)

	usersWithPassword := strings.Split(users, ",")
	for _, userWithPassword := range usersWithPassword {
		values := strings.Split(userWithPassword, ":")
		accounts[values[0]] = values[1]
	}

	router := controller.SetupRoutes(accounts)

	server := http.Server{
		Addr:         fmt.Sprintf(":%s", serverPort),
		Handler:      router,
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}

	log.Println("Starting server at port :9000")
	log.Fatal(server.ListenAndServe())
}

func getEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
