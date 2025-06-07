# TemplateAPI.OnionArchitecture.Golang
Go言語で構築したAPIサーバーのテンプレートです。  
MySQLとの接続を想定し、`air` を使ったホットリロード対応済み。  
Clean Architecture（Onion Architecture）を意識したフォルダ構成で開発しています。

## 開発環境
- Go 1.24.4
- MySQL 8.0 (Docker利用推奨)
- air v1.62.0 （Go用ホットリロードツール）
- Windows / Linux対応 (Windowsは`air.toml`で`.exe`設定済み)

## セットアップ
### 1. Go モジュールの依存解決
```bash
cd API
go mod download
```
### 2. MySQL コンテナ起動 (Docker利用時)
```sh
docker-compose up -d db
```

### 3. air のインストール
※`$GOPATH/bin` または `$HOME/go/bin` にインストールされるので、PATHを通してください。
```
go install github.com/air-verse/air@latest
```

## air の設定 (ホットリロード用)
`API/.air.toml` の設定例：
```toml
root = "."
tmp_dir = "tmp"

[build]
cmd = "go build -o ./tmp/main.exe ."
bin = "tmp/main.exe"
include_ext = ["go", "tpl", "tmpl", "html"]
exclude_dir = ["vendor"]
exclude_file = ["*_test.go"]
delay = 1000
log = "stdout"

[log]
color = "true"
time = "true"

[run]
watcher = "fsnotify"
```

## サーバー起動
```bash
cd API
air
```
air が起動し、ファイル変更を監視しつつAPIサーバーを実行します。

`http://localhost:8080/health/api` で動作確認可能です。

## DB接続設定
ソースコード内 `API/config/database_config.go` の以下の部分を適宜変更してください。
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
