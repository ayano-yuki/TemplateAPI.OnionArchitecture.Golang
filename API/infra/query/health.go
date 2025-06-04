package query

import (
	"database/sql"

	"templateapi/onionarchitecture/golang/domain/object"
)

type HealthQuery struct {
	DB *sql.DB
}

func NewHealthQuery(db *sql.DB) *HealthQuery {
	return &HealthQuery{DB: db}
}

func (q *HealthQuery) GetTestTexts() ([]object.Test, error) {
	rows, err := q.DB.Query("SELECT text FROM test")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []object.Test
	for rows.Next() {
		var text string
		if err := rows.Scan(&text); err != nil {
			return nil, err
		}
		results = append(results, object.Test{Text: text})
	}
	return results, nil
}
