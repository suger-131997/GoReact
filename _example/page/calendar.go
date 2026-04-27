package page

import (
	"context"
	"goreact/pkgs"
	"net/http"
)

type CalendarProps struct{}

func NewCalendarHandler() *pkgs.PageHandler[CalendarProps] {
	return pkgs.NewPageHandler(pkgs.PageHandlerArgs[CalendarProps]{
		EntryPoint: "page/calendar.tsx",
		HandleFunc: func(r *http.Request, render func(ctx context.Context, props CalendarProps)) {
			render(r.Context(), CalendarProps{})
		},
	})
}
