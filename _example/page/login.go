package page

import (
	"context"
	"goreact/pkgs"
	"net/http"
)

type LoginProps struct{}

func NewLoginHandler() *pkgs.PageHandler[LoginProps] {
	return pkgs.NewPageHandler(pkgs.PageHandlerArgs[LoginProps]{
		EntryPoint: "page/login.tsx",
		HandleFunc: func(r *http.Request, render func(ctx context.Context, props LoginProps)) {
			render(r.Context(), LoginProps{})
		},
	})
}
