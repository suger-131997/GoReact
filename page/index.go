package page

import (
	"context"
	"goreact/internal/contextutil"
	"goreact/internal/handler"
	"net/http"
)

type IndexProps struct {
	Name string `json:"name"`
}

func NewIndexHandler() *handler.PageHandler[IndexProps] {
	return handler.NewPageHandler(handler.PageHandlerArgs[IndexProps]{
		EntryPoint: "page/index.tsx",
		HandleFunc: func(r *http.Request, render func(ctx context.Context, props IndexProps)) {
			ctx := contextutil.WithTitle(r.Context(), "Home")

			render(ctx, IndexProps{Name: "suger"})
		},
	})
}
