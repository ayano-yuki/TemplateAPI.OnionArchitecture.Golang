package command

import (
	"database/sql"
)

type HealthCommand struct {
	DB *sql.DB
}

func NewHealthCommand(db *sql.DB) *HealthCommand {
	return &HealthCommand{DB: db}
}

func (c *HealthCommand) InsertTestText(text string) error {
	_, err := c.DB.Exec("INSERT INTO test (text) VALUES (?)", text)
	return err
}
