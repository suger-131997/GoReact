package page

import (
	"context"
	"goreact/pkgs"
	"net/http"
)

type AboutProps struct {
	Count int `json:"count"`
}

func NewAboutHandler() *pkgs.PageHandler[AboutProps] {
	count := 0

	return pkgs.NewPageHandler(pkgs.PageHandlerArgs[AboutProps]{
		EntryPoint: "page/about.tsx",
		HandleFunc: func(r *http.Request, render func(ctx context.Context, props AboutProps)) {
			count++

			render(r.Context(), AboutProps{Count: count})
		},
	})
}
