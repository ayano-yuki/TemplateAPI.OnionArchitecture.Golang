# `server/` フォルダーについて

## 目的

`server/` フォルダーは、アプリケーションの**サーバー起動および運用に関わる設定・処理**を担当します。  
WebサーバーやAPIサーバーの立ち上げ、ミドルウェア設定、サーバー固有の初期化などを管理します。

---

## 主な役割・内容

- **サーバー起動処理**  
  - HTTPサーバーやRPCサーバーの起動コードを配置します。

- **ミドルウェア設定**  
  - 認証、ログ出力、リクエスト処理などのミドルウェアの設定を管理します。

- **サーバー固有の初期化処理**  
  - ポート設定、セキュリティ設定、SSL/TLS設定などの初期化処理を含みます。

- **デプロイ・運用補助**  
  - サーバー監視やヘルスチェックエンドポイントなど運用支援に関わるコードを含む場合があります。

---

## 注意点

- ビジネスロジックは含めず、サーバーの起動および動作環境に特化してください。  
- 他層の設定やロジックとは明確に切り離して管理します。

---

## 例（イメージ）

```sh
server/
├── server           # サーバー起動スクリプト
├── middleware       # ミドルウェア設定
```