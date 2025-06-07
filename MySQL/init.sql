-- データベースがなければ作成
CREATE DATABASE IF NOT EXISTS appDB;

-- 使用するデータベースを選択
USE appDB;

-- テーブルがなければ作成
CREATE TABLE IF NOT EXISTS test (
    id INT AUTO_INCREMENT PRIMARY KEY,
    text TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
