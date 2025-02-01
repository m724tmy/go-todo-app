package models

type Todo struct {
	ID  int
	Task string
	Done bool
}

var Todos = []Todo{}

var nextID = 1

// NewTodo: 新しいTodoアイテムを作成して、データストアに追加
func NewTodo(task string) Todo {
	todo := Todo{
		ID:   nextID,
		Task: task,
		Done: false,
	}
	nextID++
	Todos = append(Todos, todo)
	return todo
}

// DeleteTodo: 指定したIDのTodoを削除
func DeleteTodo(id int) {
	for i, todo := range Todos {
		if todo.ID == id {
			//スライスから削除
			Todos = append(Todos[:i], Todos[i+1:]...)
			break
		}
	}
}
