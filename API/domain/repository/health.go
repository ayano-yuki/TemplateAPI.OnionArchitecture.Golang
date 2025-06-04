package repository

type HealthRepository interface {
	GetTestTexts() ([]string, error)
	InsertTestText(text string) error
}