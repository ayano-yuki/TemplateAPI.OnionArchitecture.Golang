package object

import "errors"

// TestEntity はtestテーブルのビジネスモデル例
type TestEntity struct {
	ID   int64
	Text string
}

// NewTestEntity はTestEntityのファクトリ関数。バリデーションもここで行う。
func NewTestEntity(id int64, text string) (*TestEntity, error) {
	if text == "" {
		return nil, errors.New("text must not be empty")
	}
	return &TestEntity{ID: id, Text: text}, nil
}

// UpdateText はTextの更新メソッド。業務ルールもここで実装可能
func (e *TestEntity) UpdateText(text string) error {
	if text == "" {
		return errors.New("text must not be empty")
	}
	e.Text = text
	return nil
}
