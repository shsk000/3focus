package main

import (
	"log"
	"os"

	"github.com/shsk/3focus/internal/handler/http"
	"github.com/shsk/3focus/internal/repository/postgres"
	"github.com/shsk/3focus/internal/usecase"
	"github.com/shsk/3focus/pkg/database"
)

func main() {
	// データベース設定
	dbConfig := database.Config{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "3focus"),
		Password: getEnv("DB_PASSWORD", "3focus_dev"),
		DBName:   getEnv("DB_NAME", "3focus_db"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}

	// データベース接続
	db, err := database.NewPostgresDB(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// 依存性注入（Dependency Injection）
	// Repository層の初期化
	bookRepo := postgres.NewBookRepository(db)

	// Usecase層の初期化（リポジトリを注入）
	bookUsecase := usecase.NewBookUsecase(bookRepo)

	// Handler層の初期化（ユースケースを注入）
	bookHandler := http.NewBookHandler(bookUsecase)

	// ルーター設定
	router := http.SetupRouter(bookHandler)

	// サーバー起動
	port := getEnv("PORT", "8080")
	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// getEnv 環境変数取得（デフォルト値付き）
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
