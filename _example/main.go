package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type PageData struct {
	Dev        bool
	ViteServer string
	Scripts    []string
	Styles     []string
}

func main() {
	viteServer := "http://localhost:5173"

	tmpl := template.Must(template.New("index").Parse(htmlTemplate))

	mux := http.NewServeMux()

	assets := http.FileServerFS(os.DirFS("."))

	mux.Handle("/assets/", assets)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" || r.URL.Path == "/index.html" {
			data := PageData{
				Dev:        true,
				ViteServer: viteServer,
			}

			w.Header().Set("Content-Type", "text/html")
			tmpl.Execute(w, data)

			return
		}

		http.ServeFileFS(w, r, os.DirFS("./public"), filepath.Base(r.URL.Path))
	})

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
    <script type="module" src="{{ .ViteServer }}/page/main.tsx"></script>
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
