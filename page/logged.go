package page

import (
	"context"
	"goreact/internal/handler"
	"goreact/internal/mark"
	"net/http"
)

type LoggedProps struct {
	Message string `json:"message"`
}

var _ mark.Props[LoggedProps]

func NewLoggedHandler() *handler.PageHandler[LoggedProps] {
	return handler.NewPageHandler(handler.PageHandlerArgs[LoggedProps]{
		EntryPoint: "page/logged.tsx",
		HandleFunc: func(r *http.Request, render func(ctx context.Context, props LoggedProps)) {
			render(r.Context(), LoggedProps{
				Message: "This is a protected page.",
			})
		},
	})
}
