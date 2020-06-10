package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/koloo91/release-bingo/model"
	"github.com/lib/pq"
	"time"
)

var (
	database *sql.DB

	insertEntryStatement      *sql.Stmt
	getAllEntriesStatement    *sql.Stmt
	getEntryStatement         *sql.Stmt
	updateEntryStatement      *sql.Stmt
	deleteEntryStatement      *sql.Stmt
	deleteAllEntriesStatement *sql.Stmt

	ErrEntryNotFound = fmt.Errorf("entry not found")
)

func SetDatabase(db *sql.DB) {
	database = db

	insertEntryStatement, _ = database.Prepare("INSERT INTO entries(id, text, checked, created, updated) VALUES ($1, $2, $3, $4, $5);")
	getAllEntriesStatement, _ = database.Prepare("SELECT id, text, checked, created, updated FROM entries;")
	getEntryStatement, _ = database.Prepare("SELECT id, text, checked, created, updated FROM entries WHERE id = $1;")
	updateEntryStatement, _ = database.Prepare("UPDATE entries SET text=$2, checked=$3, updated=$4 WHERE id = $1;")
	deleteEntryStatement, _ = database.Prepare("DELETE FROM entries WHERE id = $1;")
	deleteAllEntriesStatement, _ = database.Prepare("DELETE FROM entries;")
}

func SaveEntry(ctx context.Context, entry *model.Entry) error {
	_, err := insertEntryStatement.ExecContext(ctx, entry.Id, entry.Text, entry.Checked, entry.Created, entry.Updated)
	if err, ok := err.(*pq.Error); ok {
		return errors.New(err.Message)
	}
	return err
}

func GetAllEntries(ctx context.Context) ([]*model.Entry, error) {
	rows, err := getAllEntriesStatement.QueryContext(ctx)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			return nil, errors.New(err.Message)
		}
		return nil, err
	}

	result := make([]*model.Entry, 0)
	var id, text string
	var checked bool
	var created, updated time.Time

	for rows.Next() {
		if err := rows.Scan(&id, &text, &checked, &created, &updated); err != nil {
			if err, ok := err.(*pq.Error); ok {
				return nil, errors.New(err.Message)
			}
			return nil, err
		}

		result = append(result, &model.Entry{
			Id:      id,
			Text:    text,
			Checked: checked,
			Created: created,
			Updated: updated,
		})
	}

	return result, nil
}

func GetEntry(ctx context.Context, entryId string) (*model.Entry, error) {
	row := getEntryStatement.QueryRowContext(ctx, entryId)

	var id, text string
	var checked bool
	var created, updated time.Time

	if err := row.Scan(&id, &text, &checked, &created, &updated); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrEntryNotFound
		}
		if err, ok := err.(*pq.Error); ok {
			return nil, errors.New(err.Message)
		}
		return nil, err
	}

	return &model.Entry{
		Id:      id,
		Text:    text,
		Checked: checked,
		Created: created,
		Updated: updated,
	}, nil
}

func UpdateEntry(ctx context.Context, id, text string, checked bool) error {
	if _, err := updateEntryStatement.ExecContext(ctx, id, text, checked, time.Now()); err != nil {
		if err, ok := err.(*pq.Error); ok {
			return errors.New(err.Message)
		}
		return err
	}
	return nil
}

func DeleteEntry(ctx context.Context, id string) error {
	if _, err := deleteEntryStatement.ExecContext(ctx, id); err != nil {
		if err, ok := err.(*pq.Error); ok {
			return errors.New(err.Message)
		}
		return err
	}
	return nil
}

func DeleteAllEntries(ctx context.Context) error {
	if _, err := deleteAllEntriesStatement.ExecContext(ctx); err != nil {
		if err, ok := err.(*pq.Error); ok {
			return errors.New(err.Message)
		}
		return err
	}
	return nil
}
