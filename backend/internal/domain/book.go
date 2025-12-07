package domain

import (
	"errors"
	"time"
)

// Book エンティティ - ビジネスロジックの中心
type Book struct {
	ID          uint      `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Author      string    `json:"author" db:"author"`
	ISBN        string    `json:"isbn" db:"isbn"`
	PublishedAt time.Time `json:"published_at" db:"published_at"`
	IsAvailable bool      `json:"is_available" db:"is_available"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// ドメイン固有のエラー
var (
	ErrBookNotFound      = errors.New("book not found")
	ErrInvalidBookData   = errors.New("invalid book data")
	ErrBookNotAvailable  = errors.New("book is not available")
	ErrBookAlreadyExists = errors.New("book already exists")
)

// ビジネスロジック: 本を貸出可能にする
func (b *Book) MakeAvailable() error {
	if b.IsAvailable {
		return errors.New("book is already available")
	}
	b.IsAvailable = true
	b.UpdatedAt = time.Now()
	return nil
}

// ビジネスロジック: 本を貸出中にする
func (b *Book) Checkout() error {
	if !b.IsAvailable {
		return ErrBookNotAvailable
	}
	b.IsAvailable = false
	b.UpdatedAt = time.Now()
	return nil
}

// バリデーション
func (b *Book) Validate() error {
	if b.Title == "" {
		return errors.New("title is required")
	}
	if b.Author == "" {
		return errors.New("author is required")
	}
	if b.ISBN == "" {
		return errors.New("ISBN is required")
	}
	return nil
}
