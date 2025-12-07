package usecase

import (
	"context"

	"github.com/shsk/3focus/internal/domain"
	"github.com/shsk/3focus/internal/repository"
)

// BookUsecase ビジネスロジックを実装
type BookUsecase struct {
	bookRepo repository.BookRepository
}

// NewBookUsecase コンストラクタ
func NewBookUsecase(bookRepo repository.BookRepository) *BookUsecase {
	return &BookUsecase{
		bookRepo: bookRepo,
	}
}

// CreateBook 本を作成
func (u *BookUsecase) CreateBook(ctx context.Context, book *domain.Book) error {
	// バリデーション
	if err := book.Validate(); err != nil {
		return err
	}

	// ISBNの重複チェック（ビジネスルール）
	existingBook, err := u.bookRepo.GetByISBN(ctx, book.ISBN)
	if err != nil && err != domain.ErrBookNotFound {
		return err
	}
	if existingBook != nil {
		return domain.ErrBookAlreadyExists
	}

	// デフォルト値設定
	book.IsAvailable = true

	return u.bookRepo.Create(ctx, book)
}

// GetBook IDで本を取得
func (u *BookUsecase) GetBook(ctx context.Context, id uint) (*domain.Book, error) {
	return u.bookRepo.GetByID(ctx, id)
}

// GetAllBooks すべての本を取得
func (u *BookUsecase) GetAllBooks(ctx context.Context) ([]*domain.Book, error) {
	return u.bookRepo.GetAll(ctx)
}

// UpdateBook 本を更新
func (u *BookUsecase) UpdateBook(ctx context.Context, book *domain.Book) error {
	// バリデーション
	if err := book.Validate(); err != nil {
		return err
	}

	// 存在確認
	existing, err := u.bookRepo.GetByID(ctx, book.ID)
	if err != nil {
		return err
	}

	// ISBNが変更された場合、重複チェック
	if existing.ISBN != book.ISBN {
		duplicateBook, err := u.bookRepo.GetByISBN(ctx, book.ISBN)
		if err != nil && err != domain.ErrBookNotFound {
			return err
		}
		if duplicateBook != nil && duplicateBook.ID != book.ID {
			return domain.ErrBookAlreadyExists
		}
	}

	return u.bookRepo.Update(ctx, book)
}

// DeleteBook 本を削除
func (u *BookUsecase) DeleteBook(ctx context.Context, id uint) error {
	// 存在確認
	_, err := u.bookRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	return u.bookRepo.Delete(ctx, id)
}

// CheckoutBook 本を貸出
func (u *BookUsecase) CheckoutBook(ctx context.Context, id uint) error {
	book, err := u.bookRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	// ビジネスロジック: 本を貸出中にする
	if err := book.Checkout(); err != nil {
		return err
	}

	return u.bookRepo.Update(ctx, book)
}

// ReturnBook 本を返却
func (u *BookUsecase) ReturnBook(ctx context.Context, id uint) error {
	book, err := u.bookRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	// ビジネスロジック: 本を貸出可能にする
	if err := book.MakeAvailable(); err != nil {
		return err
	}

	return u.bookRepo.Update(ctx, book)
}
