package service

import (
	"context"

	"API/domain/object"
	"API/domain/repository"
)

// HealthService は複雑なドメインロジックを実装するドメインサービスの例
type HealthService struct {
	repo repository.HealthRepository
}

func NewHealthService(repo repository.HealthRepository) *HealthService {
	return &HealthService{repo: repo}
}

// CreateTest は新規TestEntity作成と保存のドメインロジック例
func (s *HealthService) CreateTest(ctx context.Context, text string) (*object.TestEntity, error) {
	test, err := object.NewTestEntity(0, text) // IDはDB挿入時に設定される想定
	if err != nil {
		return nil, err
	}

	if err := s.repo.InsertTest(ctx, test); err != nil {
		return nil, err
	}

	return test, nil
}
