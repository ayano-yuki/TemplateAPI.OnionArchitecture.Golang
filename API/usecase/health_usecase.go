package usecase

import (
	"context"
	"errors"
)

// HealthUsecase はヘルスチェック関連のユースケースを提供します。
type HealthUsecase struct {
	repo HealthRepository
}

// HealthRepository はDBアクセスの抽象インターフェースです。
// infra層で実装します。
type HealthRepository interface {
	GetTestTexts(ctx context.Context) ([]string, error)
	InsertTestText(ctx context.Context, text string) error
}

// NewHealthUsecase はHealthUsecaseのコンストラクタです。
func NewHealthUsecase(repo HealthRepository) *HealthUsecase {
	return &HealthUsecase{repo: repo}
}

// GetTestTexts はDBからtestテーブルのtextカラムを取得します。
func (u *HealthUsecase) GetTestTexts(ctx context.Context) ([]string, error) {
	texts, err := u.repo.GetTestTexts(ctx)
	if err != nil {
		return nil, err
	}
	return texts, nil
}

// InsertTestText はtestテーブルにtextを挿入します。
func (u *HealthUsecase) InsertTestText(ctx context.Context, text string) error {
	if text == "" {
		return errors.New("text cannot be empty")
	}
	return u.repo.InsertTestText(ctx, text)
}
