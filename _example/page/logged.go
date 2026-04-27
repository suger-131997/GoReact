package page

import (
	"context"
	"goreact/pkgs"
	"net/http"
)

type LoggedProps struct {
	Message string `json:"message"`
}

func NewLoggedHandler() *pkgs.PageHandler[LoggedProps] {
	return pkgs.NewPageHandler(pkgs.PageHandlerArgs[LoggedProps]{
		EntryPoint: "page/logged.tsx",
		HandleFunc: func(r *http.Request, render func(ctx context.Context, props LoggedProps)) {
			render(r.Context(), LoggedProps{
				Message: "This is a protected page.",
			})
		},
	})
}
