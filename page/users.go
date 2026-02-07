package page

import (
	"context"
	"fmt"
	"goreact/internal/handler"
	"net/http"
	"strconv"
)

type UsersProps struct {
	Users      []User `json:"users"`
	TotalCount int    `json:"totalCount"`
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func NewUsersHandler() *handler.PageHandler[UsersProps] {
	return handler.NewPageHandler(handler.PageHandlerArgs[UsersProps]{
		EntryPoint: "page/users.tsx",
		HandleFunc: func(r *http.Request, render func(ctx context.Context, props UsersProps)) {
			pageStr := r.URL.Query().Get("page")
			pageSizeStr := r.URL.Query().Get("pageSize")

			page, err := strconv.Atoi(pageStr)
			if err != nil {
				page = 0
			}

			pageSize, err := strconv.Atoi(pageSizeStr)
			if err != nil {
				pageSize = 10
			}

			totalCount := 100
			users := make([]User, 0, pageSize)

			start := page * pageSize
			end := start + pageSize
			if end > totalCount {
				end = totalCount
			}

			for i := start; i < end; i++ {
				id := i + 1
				users = append(users, User{
					ID:    id,
					Name:  fmt.Sprintf("User %d", id),
					Email: fmt.Sprintf("user%d@example.com", id),
					Role:  []string{"Admin", "Editor", "Viewer"}[id%3],
				})
			}

			render(r.Context(), UsersProps{
				Users:      users,
				TotalCount: totalCount,
			})
		},
	})
}
