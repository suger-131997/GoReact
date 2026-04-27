package page

import (
	"context"
	"goreact/internal/contextutil"
	"goreact/pkgs"
	"net/http"
)

type NotFoundProps struct {
	Path string `json:"path"`
}

func NewNotFoundHandler() *pkgs.PageHandler[NotFoundProps] {
	return pkgs.NewPageHandler(pkgs.PageHandlerArgs[NotFoundProps]{
		EntryPoint: "page/notfound.tsx",
		HandleFunc: func(r *http.Request, render func(ctx context.Context, props NotFoundProps)) {
			ctx := contextutil.WithStateCode(r.Context(), http.StatusNotFound)
			render(ctx, NotFoundProps{Path: r.URL.Path})
		},
	})
}
