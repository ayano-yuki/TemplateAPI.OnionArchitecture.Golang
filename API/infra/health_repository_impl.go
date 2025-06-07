package infra

import (
	"context"
	"database/sql"
)

// HealthRepositoryImpl は usecase.HealthRepository を実装します。
type HealthRepositoryImpl struct {
	db *sql.DB
}

func NewHealthRepositoryImpl(db *sql.DB) *HealthRepositoryImpl {
	return &HealthRepositoryImpl{db: db}
}

// GetTestTexts は test テーブルから text カラムをすべて取得します。
func (r *HealthRepositoryImpl) GetTestTexts(ctx context.Context) ([]string, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT text FROM test")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var texts []string
	for rows.Next() {
		var text string
		if err := rows.Scan(&text); err != nil {
			return nil, err
		}
		texts = append(texts, text)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return texts, nil
}

// InsertTestText は test テーブルに text カラムの値を挿入します。
func (r *HealthRepositoryImpl) InsertTestText(ctx context.Context, text string) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO test (text) VALUES (?)", text)
	return err
}
