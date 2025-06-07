package repository

import (
	"context"

	"API/domain/object"
)

// HealthRepository はドメインの永続化インターフェース
type HealthRepository interface {
	GetAllTests(ctx context.Context) ([]*object.TestEntity, error)
	InsertTest(ctx context.Context, test *object.TestEntity) error
}
