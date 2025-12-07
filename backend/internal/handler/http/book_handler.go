package http

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shsk/3focus/internal/domain"
	"github.com/shsk/3focus/internal/usecase"
)

// BookHandler HTTPハンドラー
type BookHandler struct {
	bookUsecase *usecase.BookUsecase
}

// NewBookHandler コンストラクタ
func NewBookHandler(bookUsecase *usecase.BookUsecase) *BookHandler {
	return &BookHandler{
		bookUsecase: bookUsecase,
	}
}

// CreateBookRequest リクエスト型
type CreateBookRequest struct {
	Title       string `json:"title" binding:"required"`
	Author      string `json:"author" binding:"required"`
	ISBN        string `json:"isbn" binding:"required"`
	PublishedAt string `json:"published_at" binding:"required"`
}

// UpdateBookRequest 更新リクエスト型
type UpdateBookRequest struct {
	Title       string `json:"title" binding:"required"`
	Author      string `json:"author" binding:"required"`
	ISBN        string `json:"isbn" binding:"required"`
	PublishedAt string `json:"published_at" binding:"required"`
	IsAvailable bool   `json:"is_available"`
}

// CreateBook 本を作成
// POST /api/v1/books
func (h *BookHandler) CreateBook(c *gin.Context) {
	var req CreateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 日付パース
	publishedAt, err := time.Parse("2006-01-02", req.PublishedAt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date format"})
		return
	}

	book := &domain.Book{
		Title:       req.Title,
		Author:      req.Author,
		ISBN:        req.ISBN,
		PublishedAt: publishedAt,
	}

	if err := h.bookUsecase.CreateBook(c.Request.Context(), book); err != nil {
		if err == domain.ErrBookAlreadyExists {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, book)
}

// GetBook 本を取得
// GET /api/v1/books/:id
func (h *BookHandler) GetBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	book, err := h.bookUsecase.GetBook(c.Request.Context(), uint(id))
	if err != nil {
		if err == domain.ErrBookNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

// GetAllBooks すべての本を取得
// GET /api/v1/books
func (h *BookHandler) GetAllBooks(c *gin.Context) {
	books, err := h.bookUsecase.GetAllBooks(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

// UpdateBook 本を更新
// PUT /api/v1/books/:id
func (h *BookHandler) UpdateBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	var req UpdateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	publishedAt, err := time.Parse("2006-01-02", req.PublishedAt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date format"})
		return
	}

	book := &domain.Book{
		ID:          uint(id),
		Title:       req.Title,
		Author:      req.Author,
		ISBN:        req.ISBN,
		PublishedAt: publishedAt,
		IsAvailable: req.IsAvailable,
	}

	if err := h.bookUsecase.UpdateBook(c.Request.Context(), book); err != nil {
		if err == domain.ErrBookNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if err == domain.ErrBookAlreadyExists {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

// DeleteBook 本を削除
// DELETE /api/v1/books/:id
func (h *BookHandler) DeleteBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	if err := h.bookUsecase.DeleteBook(c.Request.Context(), uint(id)); err != nil {
		if err == domain.ErrBookNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// CheckoutBook 本を貸出
// POST /api/v1/books/:id/checkout
func (h *BookHandler) CheckoutBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	if err := h.bookUsecase.CheckoutBook(c.Request.Context(), uint(id)); err != nil {
		if err == domain.ErrBookNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if err == domain.ErrBookNotAvailable {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "book checked out successfully"})
}

// ReturnBook 本を返却
// POST /api/v1/books/:id/return
func (h *BookHandler) ReturnBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	if err := h.bookUsecase.ReturnBook(c.Request.Context(), uint(id)); err != nil {
		if err == domain.ErrBookNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "book returned successfully"})
}
