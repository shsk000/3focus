package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/shsk/3focus/internal/domain"
)

// PostgresBookRepository PostgreSQL実装
type PostgresBookRepository struct {
	db *sqlx.DB
}

// NewBookRepository コンストラクタ
func NewBookRepository(db *sqlx.DB) *PostgresBookRepository {
	return &PostgresBookRepository{db: db}
}

// Create 本を作成
func (r *PostgresBookRepository) Create(ctx context.Context, book *domain.Book) error {
	query := `
		INSERT INTO books (title, author, isbn, published_at, is_available, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
		RETURNING id, created_at, updated_at
	`
	return r.db.QueryRowContext(
		ctx,
		query,
		book.Title,
		book.Author,
		book.ISBN,
		book.PublishedAt,
		book.IsAvailable,
	).Scan(&book.ID, &book.CreatedAt, &book.UpdatedAt)
}

// GetByID IDで本を取得
func (r *PostgresBookRepository) GetByID(ctx context.Context, id uint) (*domain.Book, error) {
	var book domain.Book
	query := `
		SELECT id, title, author, isbn, published_at, is_available, created_at, updated_at
		FROM books
		WHERE id = $1
	`
	err := r.db.GetContext(ctx, &book, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrBookNotFound
		}
		return nil, err
	}
	return &book, nil
}

// GetAll すべての本を取得
func (r *PostgresBookRepository) GetAll(ctx context.Context) ([]*domain.Book, error) {
	var books []*domain.Book
	query := `
		SELECT id, title, author, isbn, published_at, is_available, created_at, updated_at
		FROM books
		ORDER BY created_at DESC
	`
	err := r.db.SelectContext(ctx, &books, query)
	return books, err
}

// Update 本を更新
func (r *PostgresBookRepository) Update(ctx context.Context, book *domain.Book) error {
	query := `
		UPDATE books
		SET title = $1, author = $2, isbn = $3, published_at = $4, is_available = $5, updated_at = NOW()
		WHERE id = $6
		RETURNING updated_at
	`
	return r.db.QueryRowContext(
		ctx,
		query,
		book.Title,
		book.Author,
		book.ISBN,
		book.PublishedAt,
		book.IsAvailable,
		book.ID,
	).Scan(&book.UpdatedAt)
}

// Delete 本を削除
func (r *PostgresBookRepository) Delete(ctx context.Context, id uint) error {
	query := `DELETE FROM books WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return domain.ErrBookNotFound
	}

	return nil
}

// GetByISBN ISBNで本を取得
func (r *PostgresBookRepository) GetByISBN(ctx context.Context, isbn string) (*domain.Book, error) {
	var book domain.Book
	query := `
		SELECT id, title, author, isbn, published_at, is_available, created_at, updated_at
		FROM books
		WHERE isbn = $1
	`
	err := r.db.GetContext(ctx, &book, query, isbn)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrBookNotFound
		}
		return nil, err
	}
	return &book, nil
}
