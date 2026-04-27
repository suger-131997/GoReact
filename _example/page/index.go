package page

import (
	"context"
	"goreact/pkgs"
	"net/http"
)

type IndexProps struct {
	Name string `json:"name"`
}

func NewIndexHandler() *pkgs.PageHandler[IndexProps] {
	return pkgs.NewPageHandler[IndexProps](pkgs.PageHandlerArgs[IndexProps]{
		EntryPoint: "page/index.tsx",
		HandleFunc: func(r *http.Request, render func(ctx context.Context, props IndexProps)) {
			render(r.Context(), IndexProps{
				Name: "suger",
			})
		},
	})
}
