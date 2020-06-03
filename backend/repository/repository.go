package repository

import (
	"context"
	"database/sql"
	"github.com/koloo91/release-bingo/model"
	"time"
)

var (
	database *sql.DB

	insertStatement     *sql.Stmt
	getAllStatement     *sql.Stmt
	updateStatement     *sql.Stmt
	deleteByIdStatement *sql.Stmt
	deleteAllStatement  *sql.Stmt
)

func SetDatabase(db *sql.DB) {
	database = db

	insertStatement, _ = database.Prepare("INSERT INTO entries(id, text, created, updated) VALUES ($1, $2, $3, $4);")
	getAllStatement, _ = database.Prepare("SELECT id, text, created, updated FROM entries;")
	updateStatement, _ = database.Prepare("UPDATE entries SET text=$2, updated=$3 WHERE id = $1;")
	deleteByIdStatement, _ = database.Prepare("DELETE FROM entries WHERE id = $1;")
	deleteAllStatement, _ = database.Prepare("DELETE FROM entries;")
}

func SaveEntry(ctx context.Context, entry *model.Entry) error {
	_, err := insertStatement.ExecContext(ctx, entry.Id, entry.Text, entry.Created, entry.Updated)
	return err
}

func GetAll(ctx context.Context) ([]*model.Entry, error) {
	rows, err := getAllStatement.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*model.Entry, 0)
	var id, text string
	var created, updated time.Time

	for rows.Next() {
		if err := rows.Scan(&id, &text, &created, &updated); err != nil {
			return nil, err
		}

		result = append(result, &model.Entry{
			Id:      id,
			Text:    text,
			Created: created,
			Updated: updated,
		})
	}

	return result, nil
}

func UpdateEntry(ctx context.Context, id string, entry *model.Entry) error {
	if _, err := updateStatement.ExecContext(ctx, id, entry.Text, entry.Updated); err != nil {
		return err
	}
	return nil
}

func DeleteById(ctx context.Context, id string) error {
	if _, err := deleteByIdStatement.ExecContext(ctx, id); err != nil {
		return err
	}
	return nil
}

func DeleteAll(ctx context.Context) error {
	if _, err := deleteAllStatement.ExecContext(ctx); err != nil {
		return err
	}
	return nil
}
