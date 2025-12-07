-- 3focus Database Schema
-- "Today's 3 Tasks Only" Application

-- ============================================
-- テーブル設計
-- ============================================

-- TODO: 必要なテーブルを設計してください
--
-- 検討すべきテーブル:
-- - users (ユーザー情報)
-- - tasks (タスク情報)
-- - task_history (過去のタスク履歴)
--
-- 検討すべきカラム:
-- - ID、作成日時、更新日時
-- - タスク: タイトル、優先度(1-3)、完了状態、日付
-- - ユーザー: 認証情報（必要に応じて）

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    auth0_user_id TEXT NOT NULL UNIQUE,
    email TEXT NOT NUL CHECK (LENGTH(email) <= 255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_users_auth0_user_id ON users(auth0_user_id);
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);

CREATE TABLE IF NOT EXISTS todos (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL FOREIGN KEY REFERENCES users(id),
    title TEXT NOT NULL CHECK (LENGTH(title) <= 100),
    description TEXT CHECK (LENGTH(description) <= 1000),
    priority SMALLINT NOT NULL CHECK (priority BETWEEN 1 AND 3),
    is_completed BOOLEAN NOT NULL DEFAULT FALSE,
    completed_at TIMESTAMPTZ,
    task_date DATE NOT NULL DEFAULT CURRENT_DATE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_todos_user_date ON todos(user_id, task_date);