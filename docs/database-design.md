# データベース設計

## 概要

3focus アプリケーションのデータベース設計ドキュメント

## 要件

application-plan.md を参照：
- 1日に3つまでのタスク
- 優先度付き（①②③）
- 完了状態の管理
- 日次でのデータ管理
- 履歴の保存

## テーブル設計

### TODO: テーブル構成を設計してください

検討ポイント：
- どのようなテーブルが必要か？
- 各テーブルにどんなカラムが必要か？
- テーブル間のリレーションは？
- インデックスは何が必要か？
- 制約（UNIQUE、NOT NULL など）は？

### todos
| カラム名 | 型 | 説明 |
| -------- | ---- | ---- |
| id | serial primary key | タスクID |
| user_id | uuid not null foreign key references users(id) | ユーザーID |
| title | text not null check (length(title) <= 100) | タスク名 |
| description | text check (length(description) <= 1000) | タスク説明 |
| priority | smallint not null check (priority between 1 and 3) | 優先度 |
| is_completed | boolean not null default false | 完了状態 |
| completed_at | timestamptz | 完了日時 |
| created_at | timestamptz | 作成日時 |
| updated_at | timestamptz | 更新日時 |
| task_date | date not null default current_date | タスク日付 |

CREATE INDEX idx_todos_user_date ON todos(user_id, task_date);

### users

| カラム名 | 型 | 説明 |
| -------- | ---- | ---- |
| id | uuid primary key default gen_random_uuid() | ユーザーID |
| auth0_user_id | text not null unique | Auth0のユーザーID |
| email | text not null check (length(email) <= 255) | ユーザーのメールアドレス |
| created_at | timestamptz | 作成日時 |
| updated_at | timestamptz | 更新日時 |

## ER図

```
TODO: テーブル間の関係を図示してください
```

## マイグレーション戦略

TODO: マイグレーションツールの選定と戦略

## Notes

- PostgreSQL 16を使用
- タイムゾーン: UTCで保存、表示時に変換

