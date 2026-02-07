package server

import (
	"fmt"
	static "goreact"
	"goreact/api"
	"goreact/internal/render"
	"goreact/middleware"
	"goreact/page"
	"io/fs"
	"log"
	"net/http"
)

type JSLoader func(entryPoint string) func(entryPoint string) (string, error)

func NewRouter(renderer *render.CoreRenderer) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func() func(writer http.ResponseWriter, request *http.Request) {
		indexHandler := page.NewIndexHandler().Handler(renderer)
		notFoundHandler := page.NewNotFoundHandler().Handler(renderer)
		return func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/" {
				notFoundHandler(w, r)
				return
			}
			indexHandler(w, r)
		}
	}())

	mux.HandleFunc("/login", page.NewLoginHandler().Handler(renderer))
	mux.HandleFunc("/about", page.NewAboutHandler().Handler(renderer))
	mux.HandleFunc("/calendar", page.NewCalendarHandler().Handler(renderer))
	mux.HandleFunc("/users", page.NewUsersHandler().Handler(renderer))
	mux.Handle("/logged", middleware.AuthMiddleware(page.NewLoggedHandler().Handler(renderer)))

	mux.HandleFunc("POST /api/login", api.LoginHandler)
	mux.HandleFunc("POST /api/logout", api.LogoutHandler)

	assetsFS, err := fs.Sub(static.Assets, "assets")
	if err != nil {
		log.Fatalf("Error initializing assets file system: %v", err)
	}
	mux.Handle("/_assets/", http.StripPrefix("/_assets/",
		http.FileServer(http.FS(assetsFS)),
	))

	return mux
}

func Start(mux *http.ServeMux, port int) {
	fmt.Printf("Server is running at http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), middleware.RequestLogger(middleware.MetadataMiddleware(mux, "GoReact"))))
}
