package page

import (
	"context"
	"goreact/internal/handler"
	"net/http"
)

type TodoProps struct {
	InitialTodos []TodoItem `json:"initialTodos"`
}

type TodoItem struct {
	ID        int    `json:"id"`
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}

func NewTodoHandler() *handler.PageHandler[TodoProps] {
	return handler.NewPageHandler(handler.PageHandlerArgs[TodoProps]{
		EntryPoint: "page/todo.tsx",
		HandleFunc: func(r *http.Request, render func(ctx context.Context, props TodoProps)) {
			// 初期表示用のダミーデータ
			initialTodos := []TodoItem{
				{ID: 1, Text: "Learn Go", Completed: true},
				{ID: 2, Text: "Learn React", Completed: true},
				{ID: 3, Text: "Build a Todo App", Completed: false},
			}

			render(r.Context(), TodoProps{
				InitialTodos: initialTodos,
			})
		},
	})
}
