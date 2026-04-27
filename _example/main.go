package main

import (
	"context"
	"fmt"
	"goreact/_example/page"
	"goreact/pkgs"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type PageData struct {
	Dev           bool
	ViteServer    string
	EntryPointDir string
	EntryPoint    string
	Scripts       []string
	Styles        []string
}

func main() {
	viteServer := "http://localhost:5173"
	workdir := "tmp"

	// Cleanup workdir
	if err := os.RemoveAll(workdir); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll(workdir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	g := pkgs.NewEntryPointGenerator(workdir)

	ctx = pkgs.WithEntryPointGenerator(ctx, g)

	ctx, err := pkgs.WithRenderCreatorForDev(ctx, htmlDevTemplate, viteServer, workdir)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.Handle("/assets/", http.FileServerFS(os.DirFS(".")))

	mux.HandleFunc("/", func() func(writer http.ResponseWriter, request *http.Request) {
		indexHandler := page.NewIndexHandler().Handler(ctx)
		return func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/" || r.URL.Path == "/index.html" {
				indexHandler(w, r)
				return
			}

			http.ServeFileFS(w, r, os.DirFS("./public"), filepath.Base(r.URL.Path))
		}
	}())

	mux.HandleFunc("/app", page.NewAppHandler().Handler(ctx))

	port := ":8080"
	fmt.Printf("Server started at http://localhost%s\n", port)

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}

const htmlTemplate = `
<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <link rel="icon" type="image/svg+xml" href="/favicon.svg" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Vite + Go Integration</title>
    {{ if .Dev }}
    <!-- Vite 開発環境用の設定 -->
    <script type="module">
        import RefreshRuntime from '{{ .ViteServer }}/@react-refresh'
        RefreshRuntime.injectIntoGlobalHook(window)
        window.$RefreshReg$ = () => {}
        window.$RefreshSig$ = () => (type) => type
        window.__vite_plugin_react_preamble_installed__ = true
    </script>
    <script type="module" src="{{ .ViteServer }}/@vite/client"></script>
    <script type="module" src="{{ .ViteServer }}/{{ .EntryPointDir }}/{{ .EntryPoint }}"></script>
    {{ else }}
    <!-- 本番環境（ビルド済みアセット） -->
    {{ range .Styles }}
    <link rel="stylesheet" href="{{ . }}">
    {{ end }}
    {{ range .Scripts }}
    <script type="module" src="{{ . }}"></script>
    {{ end }}
    {{ end }}
</head>
<body>
    <div id="root"></div>
</body>
</html>
`

const htmlDevTemplate = `
<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <link rel="icon" type="image/svg+xml" href="/favicon.svg" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Vite + Go Integration</title>
	<script>window.APP_PROPS={{ .AppProps }};</script>
    <script type="module">
        import RefreshRuntime from '{{ .ViteServer }}/@react-refresh'
        RefreshRuntime.injectIntoGlobalHook(window)
        window.$RefreshReg$ = () => {}
        window.$RefreshSig$ = () => (type) => type
        window.__vite_plugin_react_preamble_installed__ = true
    </script>
    <script type="module" src="{{ .ViteServer }}/@vite/client"></script>
    <script type="module" src="{{ .ViteServer }}/{{ .Workdir }}/{{ .EntryPoint }}"></script>
</head>
<body>
    <div id="root"></div>
</body>
</html>
`
