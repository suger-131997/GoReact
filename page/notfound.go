package page

import (
	"context"
	"goreact/internal/contextutil"
	"goreact/internal/handler"
	"net/http"
)

type NotFoundProps struct {
	Path string `json:"path"`
}

func NewNotFoundHandler() *handler.PageHandler[NotFoundProps] {
	return handler.NewPageHandler(handler.PageHandlerArgs[NotFoundProps]{
		EntryPoint: "page/notfound.tsx",
		HandleFunc: func(r *http.Request, render func(ctx context.Context, props NotFoundProps)) {
			ctx := contextutil.WithStateCode(r.Context(), http.StatusNotFound)
			render(ctx, NotFoundProps{Path: r.URL.Path})
		},
	})
}
