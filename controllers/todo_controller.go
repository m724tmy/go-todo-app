package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/m724tmy/go-todo-app/models"
)

// IndexTodo: Todoの一覧を表示
func IndexTodo(c *gin.Context) {
	c.HTML(http.StatusOK, "content", gin.H{
		"todos": models.Todos,
	})
}

//CreateTodo: POSTされたタスク内容から新規Todoを作成
func CreateTodo(c *gin.Context) {
	task := c.PostForm("task")
	if task != "" {
		models.NewTodo(task)
	}
	c.Redirect(http.StatusFound, "/")
}

// DeleteTodo: 指定されたIDのTodoを削除
func DeleteTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err == nil {
		models.DeleteTodo(id)
	}
	c.Redirect(http.StatusFound, "/")
}