package page

import (
	"context"
	"goreact/pkgs"
	"net/http"
)

type AppProps struct{}

func NewAppHandler() *pkgs.PageHandler[AppProps] {
	return pkgs.NewPageHandler[AppProps](pkgs.PageHandlerArgs[AppProps]{
		EntryPoint: "page/app.tsx",
		HandleFunc: func(r *http.Request, render func(ctx context.Context, props AppProps)) {
			render(r.Context(), AppProps{})
		},
	})
}
