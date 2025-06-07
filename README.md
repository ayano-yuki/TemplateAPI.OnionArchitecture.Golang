# TemplateAPI.OnionArchitecture.Golang
Go言語のChi × SwagerUI × オニオンアーキテクチャで作成するWeb APIのテンプレートです。DBはMySQLLとの接続を想定し、air を使ったホットリロードにも対応しています。

## 機能
- ホットリロード（[Air](https://github.com/air-verse/air)）
- オニオンアーキテクチャによる分離設計
- MySQL連携
- Swagger UIによるAPIドキュメント生成（[swaggo](https://github.com/swaggo/swag)）

## 開発環境の立ち上げ
### 前提条件
- Go 1.24+
- Docker / Docker Compose
- Air (`go install github.com/air-verse/air@latest`)
- swag (`go install github.com/swaggo/swag/cmd/swag@latest`)

### 立ち上げコマンド
1. MySQLをDockerで立ち上げる
   ```bash
	docker-compose up --build

   ```
2. APIを立ち上げる
	```bash
	# 依存取得
	go mod tidy

	# Swagger ドキュメント生成
	swag init --parseDependency --parseInternal

	# ホットリロードで起動（Air）
	# APIは http://localhost:8080 にて起動します。
	# SwaggerUIは http://localhost:8080/swagger/ にてアクセスできる。
	air
	```
## データベース接続設定
環境変数による DB 接続情報は `config/database_config.go` 直書きしている
```go
func LoadDBConfig() *DBConfig {
	return &DBConfig{
		Host:     "localhost",
		Port:     "3306",
		User:     "user",
		Password: "pass",
		Name:     "appDB",
	}
}

```

## エラー共通レスポンス構造（例）
全ハンドラでこの形式を使用し、config パッケージで集中管理しています。
```go
type ErrorResponse struct {
  Code    int    `json:"code"`
  Message string `json:"message"`
}
```

## 注意事項
- `tmp/` ディレクトリは Air によって一時生成され、.gitignore 済です。
- `air` を実行すると `docs/` にSwagger関係のファイルが生成されます（バージョン管理推奨）。