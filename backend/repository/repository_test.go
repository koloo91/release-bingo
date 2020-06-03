package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/docker/go-connections/nat"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/koloo91/release-bingo/model"
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

type RepositoryTestSuite struct {
	suite.Suite
	postgresContainer testcontainers.Container
}

func (suite *RepositoryTestSuite) SetupSuite() {
	log.Println("Setup suite")

	req := testcontainers.ContainerRequest{
		Image:        "postgres",
		ExposedPorts: []string{"5432/tcp"},
		Env:          map[string]string{"POSTGRES_PASSWORD": "Pass00"},
		WaitingFor: wait.ForSQL("5432/tcp", "postgres", func(port nat.Port) string {
			return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port.Port(), user, password, dbname)
		}),
	}

	postgresContainer, err := testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	suite.postgresContainer = postgresContainer
	if err != nil {
		log.Fatalf("Error starting postgres testcontainer: %v", err)
	}

	port, err := postgresContainer.MappedPort(context.Background(), "5432")
	if err != nil {
		log.Fatalf("Error getting mapped port: %v", err)
	}

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port.Port(), user, password, dbname)
	db, _ := sql.Open("postgres", connectionString)

	driver, _ := postgres.WithInstance(db, &postgres.Config{})
	m, _ := migrate.NewWithDatabaseInstance("file://../migrations", "postgres", driver)
	if err := m.Up(); err != nil {
		log.Printf("Error running migrations: %v", err)
	}

	SetDatabase(db)
}

func (suite *RepositoryTestSuite) SetupTest() {
	log.Println("Setup test")
	DeleteAll(context.Background())
}

func (suite *RepositoryTestSuite) TearDownSuite() {
	suite.postgresContainer.Terminate(context.Background())
}

func TestRepositorySuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}

func (suite *RepositoryTestSuite) TestSaveEntry() {
	entry := model.NewEntry("Hello")

	err := SaveEntry(context.Background(), entry)
	suite.Nil(err)

	err = SaveEntry(context.Background(), entry)
	suite.NotNil(err)
}

func (suite *RepositoryTestSuite) TestGetAll() {
	for i := 0; i < 10; i++ {
		entry := model.NewEntry(fmt.Sprintf("Hello %d", i))
		err := SaveEntry(context.Background(), entry)
		suite.Nil(err)
	}

	all, err := GetAll(context.Background())
	suite.Nil(err)
	suite.Equal(10, len(all))

	for i, entry := range all {
		suite.Equal(fmt.Sprintf("Hello %d", i), entry.Text)
	}
}
