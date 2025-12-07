package http

import (
	"github.com/gin-gonic/gin"
)

// SetupRouter ルーティング設定
func SetupRouter(bookHandler *BookHandler) *gin.Engine {
	r := gin.Default()

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// API v1
	v1 := r.Group("/api/v1")
	{
		// Books endpoints
		books := v1.Group("/books")
		{
			books.POST("", bookHandler.CreateBook)           // 本を作成
			books.GET("", bookHandler.GetAllBooks)           // すべての本を取得
			books.GET("/:id", bookHandler.GetBook)           // 本を取得
			books.PUT("/:id", bookHandler.UpdateBook)        // 本を更新
			books.DELETE("/:id", bookHandler.DeleteBook)     // 本を削除
			books.POST("/:id/checkout", bookHandler.CheckoutBook) // 貸出
			books.POST("/:id/return", bookHandler.ReturnBook)     // 返却
		}
	}

	return r
}
