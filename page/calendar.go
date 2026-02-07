package page

import (
	"context"
	"goreact/internal/handler"
	"net/http"
)

type CalendarProps struct {
}

func NewCalendarHandler() *handler.PageHandler[CalendarProps] {
	return handler.NewPageHandler(handler.PageHandlerArgs[CalendarProps]{
		EntryPoint: "page/calendar.tsx",
		HandleFunc: func(r *http.Request, render func(ctx context.Context, props CalendarProps)) {
			render(r.Context(), CalendarProps{})
		},
	})
}
