package page

import (
	"context"
	"goreact/pkgs"
	"net/http"
	"sync"
)

type AboutProps struct {
	Count int `json:"count"`
}

func NewAboutHandler() *pkgs.PageHandler[AboutProps] {
	var mu sync.Mutex
	count := 0

	return pkgs.NewPageHandler(pkgs.PageHandlerArgs[AboutProps]{
		EntryPoint: "page/about.tsx",
		HandleFunc: func(r *http.Request, render func(ctx context.Context, props AboutProps)) {
			mu.Lock()
			count++
			current := count
			mu.Unlock()

			render(r.Context(), AboutProps{Count: current})
		},
	})
}
