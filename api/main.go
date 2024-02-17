package main

import "github.com/gin-gonic/gin"

func main() {
	// Ginエンジンの初期化
	router := gin.Default()

	// ヘルスチェック用のルート
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "UP"})
	})

	// サーバーの起動
	router.Run()
}
