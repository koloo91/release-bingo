package service

import (
	"context"
	"github.com/koloo91/release-bingo/model"
	"github.com/koloo91/release-bingo/repository"
)

func CreateEntry(ctx context.Context, entry *model.Entry) (*model.Entry, error) {
	return entry, repository.SaveEntry(ctx, entry)
}

func GetEntries(ctx context.Context) ([]*model.Entry, error) {
	return repository.GetAllEntries(ctx)
}

func GetEntryById(ctx context.Context, id string) (*model.Entry, error) {
	return repository.GetEntry(ctx, id)
}

func UpdateEntry(ctx context.Context, id, text string, checked bool) (*model.Entry, error) {
	if err := repository.UpdateEntry(ctx, id, text, checked); err != nil {
		return nil, err
	}

	return repository.GetEntry(ctx, id)
}

func DeleteEntry(ctx context.Context, id string) error {
	return repository.DeleteEntry(ctx, id)
}
