package page

import (
	"context"
	"goreact/internal/handler"
	"net/http"
)

type LoginProps struct{}

func NewLoginHandler() *handler.PageHandler[LoginProps] {
	return handler.NewPageHandler(handler.PageHandlerArgs[LoginProps]{
		EntryPoint: "page/login.tsx",
		HandleFunc: func(r *http.Request, render func(ctx context.Context, props LoginProps)) {
			render(r.Context(), LoginProps{})
		},
	})
}
