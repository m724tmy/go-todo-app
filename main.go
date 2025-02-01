package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m724tmy/go-todo-app/controllers"
)

func main() {
	// Ginデフォルトのルーターを作成
	router := gin.Default()

	// template読み込み
	// layout.tmplをインクルード、index.tmpl読み込み
	router.LoadHTMLGlob("templates/*")

	// ルーティング設定
	router.GET("/", controllers.IndexTodo)
	router.POST("/todo", controllers.CreateTodo)
	// 削除はGETで実装（本番ではPOST/DELETE推奨）
	router.GET("/todo/delete/:id", controllers.DeleteTodo)

	// 静的ファイルが必要な場合は以下のように設定する
	// router.Static("/assets", "./assets")

	// サーバー起動
	if err := router.Run(":8080"); err != nil && err != http.ErrServerClosed {
		log.Fatalf("サーバー起動エラー: %v", err)
	}
}