# 3focus Backend - Clean Architecture Sample

このディレクトリには、**Clean Architecture**を実装したサンプルコード（Book管理API）が含まれています。

## ディレクトリ構成

```
backend/
├── cmd/
│   └── api/
│       └── main.go                     # エントリーポイント（依存性注入）
│
├── internal/
│   ├── domain/                         # ドメイン層（ビジネスエンティティ）
│   │   └── book.go                     # Bookエンティティ＋ビジネスロジック
│   │
│   ├── usecase/                        # ユースケース層（アプリケーションロジック）
│   │   └── book_usecase.go             # 本の操作に関するビジネスルール
│   │
│   ├── repository/                     # リポジトリ層（データアクセス）
│   │   ├── book_repository.go          # インターフェース定義
│   │   └── postgres/                   # PostgreSQL実装
│   │       └── book_repository.go
│   │
│   └── handler/                        # ハンドラー層（HTTP API）
│       └── http/
│           ├── book_handler.go         # HTTPハンドラー
│           └── router.go               # ルーティング設定
│
├── pkg/                                # 外部公開可能なユーティリティ
│   └── database/
│       └── db.go                       # DB接続管理
│
├── migrations/                         # マイグレーションファイル
│   ├── 000003_create_books_sample.up.sql
│   └── 000003_create_books_sample.down.sql
│
├── go.mod
└── go.sum
```

## 層の責務と依存関係

### 依存関係の流れ（Clean Architectureの原則）

```
Handler → Usecase → Repository Interface
                ↓
         Domain Entity
                ↑
   Repository Implementation (Postgres)
```

### 各層の役割

1. **Domain層** (`internal/domain/`)
   - ビジネスエンティティ
   - ビジネスルール
   - 他の層に依存しない（最内部）

2. **Usecase層** (`internal/usecase/`)
   - アプリケーション固有のビジネスロジック
   - Domainとリポジトリインターフェースに依存

3. **Repository層** (`internal/repository/`)
   - データアクセスのインターフェース定義
   - 実装は `postgres/` に配置（依存性逆転）

4. **Handler層** (`internal/handler/`)
   - HTTP API（配送層）
   - Usecaseに依存

## サンプルAPIエンドポイント

### 本の管理

```bash
# 本を作成
POST /api/v1/books
{
  "title": "Clean Architecture",
  "author": "Robert C. Martin",
  "isbn": "978-0134494166",
  "published_at": "2017-09-20"
}

# すべての本を取得
GET /api/v1/books

# 本を取得
GET /api/v1/books/:id

# 本を更新
PUT /api/v1/books/:id
{
  "title": "Clean Architecture (Updated)",
  "author": "Robert C. Martin",
  "isbn": "978-0134494166",
  "published_at": "2017-09-20",
  "is_available": true
}

# 本を削除
DELETE /api/v1/books/:id

# 本を貸出
POST /api/v1/books/:id/checkout

# 本を返却
POST /api/v1/books/:id/return
```

## セットアップ＆実行

### 1. データベースマイグレーション

```bash
# コンテナ起動
mise run up

# マイグレーション実行（booksテーブル作成）
docker-compose exec -T db psql -U 3focus -d 3focus_db < migrations/000003_create_books_sample.up.sql
```

### 2. 依存関係のインストール

```bash
cd backend
go mod download
```

### 3. サーバー起動（ローカル）

```bash
# 環境変数設定
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=3focus
export DB_PASSWORD=3focus_dev
export DB_NAME=3focus_db
export PORT=8080

# サーバー起動
go run cmd/api/main.go
```

### 4. Dockerで起動（推奨）

```bash
# バックエンドコンテナが既に起動している場合
docker-compose restart backend

# ログ確認
docker-compose logs -f backend
```

## 学習ポイント

このサンプルコードで学べること：

### 1. **依存性逆転の原則**
```go
// Usecaseはインターフェースに依存（具体的な実装に依存しない）
type BookUsecase struct {
    bookRepo repository.BookRepository // インターフェース
}
```

### 2. **ビジネスロジックの配置**
```go
// Domain層にビジネスルール
func (b *Book) Checkout() error {
    if !b.IsAvailable {
        return ErrBookNotAvailable
    }
    b.IsAvailable = false
    return nil
}

// Usecase層にアプリケーション固有のロジック
func (u *BookUsecase) CheckoutBook(ctx context.Context, id uint) error {
    book, _ := u.bookRepo.GetByID(ctx, id)
    return book.Checkout() // ドメインロジックを呼び出し
}
```

### 3. **テストしやすい設計**
各層をモックして独立してテスト可能：

```go
// リポジトリをモック
type MockBookRepository struct {
    books map[uint]*domain.Book
}

// Usecaseをテスト
func TestCreateBook(t *testing.T) {
    mockRepo := &MockBookRepository{books: make(map[uint]*domain.Book)}
    usecase := usecase.NewBookUsecase(mockRepo)
    // テスト実行...
}
```

### 4. **main.goでの依存性注入**
```go
// 外側から内側への依存性注入
bookRepo := postgres.NewBookRepository(db)
bookUsecase := usecase.NewBookUsecase(bookRepo)
bookHandler := http.NewBookHandler(bookUsecase)
```

## 3focus実装への適用

このサンプルを参考に、Todo/User機能を実装する際：

1. **Domain層**: `Todo`, `User` エンティティを定義
2. **Repository層**: データアクセスインターフェースを定義
3. **Usecase層**: "1日3つまで" などのビジネスルールを実装
4. **Handler層**: REST APIエンドポイントを実装

同じパターンで実装できます！

## 次のステップ

- [ ] サンプルAPIを動かして理解を深める
- [ ] Todo/User機能を同じパターンで実装
- [ ] Auth0認証をミドルウェアで追加
- [ ] テストコードを書く
