package integrationtest

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/docker/go-connections/nat"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/koloo91/release-bingo/repository"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"testing"
)

const (
	host     = "localhost"
	user     = "postgres"
	password = "Pass00"
	dbname   = "postgres"
)

type BaseTestSuite struct {
	suite.Suite
	postgresContainer testcontainers.Container
}

func (suite *BaseTestSuite) SetupSuite() {
	log.Println("Setup suite")

	postgresContainer, err := setupPostgresTestContainer()
	if err != nil {
		log.Fatalf("Error starting postgres testcontainer: %v", err)
	}

	suite.postgresContainer = postgresContainer

	port, err := postgresContainer.MappedPort(context.Background(), "5432")
	if err != nil {
		log.Fatalf("Error getting mapped port: %v", err)
	}

	connectionString := fmt.Sprintf("wsHost=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port.Port(), user, password, dbname)
	db, _ := sql.Open("postgres", connectionString)

	driver, _ := postgres.WithInstance(db, &postgres.Config{})
	m, _ := migrate.NewWithDatabaseInstance("file://../migrations", "postgres", driver)
	if err := m.Up(); err != nil {
		log.Printf("Error running migrations: %v", err)
	}

	repository.SetDatabase(db)
}

func (suite *BaseTestSuite) SetupTest() {
	log.Println("Setup test")
	repository.DeleteAllEntries(context.Background())
}

func (suite *BaseTestSuite) TearDownSuite() {
	suite.postgresContainer.Terminate(context.Background())
}

func TestRepositorySuite(t *testing.T) {
	suite.Run(t, new(BaseTestSuite))
}

func setupPostgresTestContainer() (testcontainers.Container, error) {
	req := testcontainers.ContainerRequest{
		Image:        "postgres",
		ExposedPorts: []string{"5432/tcp"},
		Env:          map[string]string{"POSTGRES_PASSWORD": "Pass00"},
		WaitingFor: wait.ForSQL("5432/tcp", "postgres", func(port nat.Port) string {
			return fmt.Sprintf("wsHost=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port.Port(), user, password, dbname)
		}),
	}

	return testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
}
