-- 3focus Database Schema
-- "Today's 3 Tasks Only" Application

-- ============================================
-- テーブル設計
-- ============================================

-- Users テーブル（Auth0認証）
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    auth0_user_id TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL CHECK (LENGTH(email) <= 255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_users_auth0_user_id ON users(auth0_user_id);

-- Todos テーブル（タスク管理）
CREATE TABLE IF NOT EXISTS todos (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id),
    title TEXT NOT NULL CHECK (LENGTH(title) <= 100),
    priority SMALLINT NOT NULL CHECK (priority BETWEEN 1 AND 3),
    is_completed BOOLEAN NOT NULL DEFAULT FALSE,
    completed_at TIMESTAMPTZ,
    task_date DATE NOT NULL DEFAULT CURRENT_DATE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_todos_user_date ON todos(user_id, task_date);

-- ============================================
-- 注意事項
-- ============================================
--
-- アプリケーション側で実装すべき機能:
-- 1. updated_at の自動更新
--    - GORM の autoUpdateTime を使用
--    - または手動で time.Now() を設定
--
-- 2. 1日3つまでの制約
--    - タスク作成前にカウントチェック
--    - 3つ以上の場合はエラーを返す
--
