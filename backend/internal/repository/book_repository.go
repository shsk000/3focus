package repository

import (
	"context"

	"github.com/shsk/3focus/internal/domain"
)

// BookRepository インターフェース
// このインターフェースはドメイン層で使用される
// 実装は postgres パッケージで行う（依存性逆転の原則）
type BookRepository interface {
	Create(ctx context.Context, book *domain.Book) error
	GetByID(ctx context.Context, id uint) (*domain.Book, error)
	GetAll(ctx context.Context) ([]*domain.Book, error)
	Update(ctx context.Context, book *domain.Book) error
	Delete(ctx context.Context, id uint) error
	GetByISBN(ctx context.Context, isbn string) (*domain.Book, error)
}
