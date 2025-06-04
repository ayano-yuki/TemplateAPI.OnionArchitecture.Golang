package usecase

import "templateapi/onionarchitecture/golang/domain/object"

type HealthRepository interface {
	GetTestTexts() ([]object.Test, error)
	InsertTestText(text string) error
}

type HealthUsecase struct {
	Repo HealthRepository
}

func NewHealthUsecase(repo HealthRepository) *HealthUsecase {
	return &HealthUsecase{Repo: repo}
}

func (u *HealthUsecase) GetTestTexts() ([]object.Test, error) {
	return u.Repo.GetTestTexts()
}

func (u *HealthUsecase) InsertText() error {
	return u.Repo.InsertTestText("文字列1")
}
