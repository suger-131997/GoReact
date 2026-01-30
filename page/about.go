package page

import (
	"context"
	"goreact/internal/handler"
	"goreact/internal/mark"
	"net/http"
)

type AboutProps struct {
	Count int `json:"count"`
}

var _ mark.Props[AboutProps]

func NewAboutHandler() *handler.PageHandler[AboutProps] {
	count := 0

	return handler.NewPageHandler(handler.PageHandlerArgs[AboutProps]{
		EntryPoint: "page/about.tsx",
		HandleFunc: func(r *http.Request, render func(ctx context.Context, props AboutProps)) {
			count++

			render(r.Context(), AboutProps{Count: count})
		},
	})
}
