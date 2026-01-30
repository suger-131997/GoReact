package page

import (
	"context"
	"goreact/internal/contextutil"
	"goreact/internal/handler"
	"goreact/internal/mark"
	"net/http"
)

type NotFoundProps struct {
	Path string `json:"path"`
}

var _ mark.Props[NotFoundProps]

func NewNotFoundHandler() *handler.PageHandler[NotFoundProps] {
	return handler.NewPageHandler(handler.PageHandlerArgs[NotFoundProps]{
		EntryPoint: "page/notfound.tsx",
		HandleFunc: func(r *http.Request, render func(ctx context.Context, props NotFoundProps)) {
			ctx := contextutil.WithStateCode(r.Context(), http.StatusNotFound)
			render(ctx, NotFoundProps{Path: r.URL.Path})
		},
	})
}
