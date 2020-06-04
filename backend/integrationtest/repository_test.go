package integrationtest

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/koloo91/release-bingo/model"
	"github.com/koloo91/release-bingo/repository"
)

func (suite *BaseTestSuite) TestSaveEntry() {
	entry := model.NewEntry("Hello")

	suite.Nil(repository.SaveEntry(context.Background(), entry))
	suite.NotNil(repository.SaveEntry(context.Background(), entry))
}

func (suite *BaseTestSuite) TestGetAll() {
	for i := 0; i < 10; i++ {
		entry := model.NewEntry(fmt.Sprintf("Hello %d", i))
		suite.Nil(repository.SaveEntry(context.Background(), entry))
	}

	all, err := repository.GetAllEntries(context.Background())
	suite.Nil(err)
	suite.Equal(10, len(all))

	for i, entry := range all {
		suite.Equal(fmt.Sprintf("Hello %d", i), entry.Text)
	}
}

func (suite *BaseTestSuite) TestGetEntry() {
	entry := model.NewEntry("Hello")
	suite.Nil(repository.SaveEntry(context.Background(), entry))

	entryFromDb, err := repository.GetEntry(context.Background(), entry.Id)
	suite.Nil(err)

	suite.Equal(entry.Id, entryFromDb.Id)
	suite.Equal(entry.Text, entryFromDb.Text)
	suite.NotNil(entryFromDb.Created)
	suite.NotNil(entryFromDb.Updated)

	_, err = repository.GetEntry(context.Background(), uuid.New().String())
	suite.NotNil(err)
	suite.Equal(repository.ErrEntryNotFound, err)

	_, err = repository.GetEntry(context.Background(), "Foo")
	suite.NotNil(err)
	suite.Equal("invalid input syntax for type uuid: \"Foo\"", err.Error())
}

func (suite *BaseTestSuite) TestUpdateEntry() {
	entry := model.NewEntry("Hello")
	suite.Nil(repository.SaveEntry(context.Background(), entry))

	suite.Nil(repository.UpdateEntry(context.Background(), entry.Id, "Hello world"))

	entryFromDb, err := repository.GetEntry(context.Background(), entry.Id)
	suite.Nil(err)

	suite.Equal(entry.Id, entryFromDb.Id)
	suite.Equal("Hello world", entryFromDb.Text)

	err = repository.UpdateEntry(context.Background(), uuid.New().String(), "Hello world")
	suite.Nil(err)
}

func (suite *BaseTestSuite) TestDeleteEntry() {
	entry := model.NewEntry("Hello")
	suite.Nil(repository.SaveEntry(context.Background(), entry))

	suite.Nil(repository.DeleteEntry(context.Background(), entry.Id))
	_, err := repository.GetEntry(context.Background(), entry.Id)
	suite.Equal(repository.ErrEntryNotFound, err)

	suite.Nil(repository.DeleteEntry(context.Background(), entry.Id))
}
