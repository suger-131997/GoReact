package main

import (
	"fmt"
	static "goreact"
	"goreact/internal/render"
	"goreact/server"
	"io"
	"path/filepath"
	"strings"
)

func main() {
	jsf := func(entryPoint string) func() (string, error) {
		jsPath := strings.TrimSuffix(entryPoint, filepath.Ext(entryPoint)) + ".js"

		f, err := static.PageFS.Open(filepath.Join(".output", jsPath))
		if err != nil {
			return func() (string, error) {
				return "", fmt.Errorf("failed to open client bundle %s: %w", jsPath, err)
			}
		}
		defer f.Close()

		js, err := io.ReadAll(f)

		if err != nil {
			return func() (string, error) {
				return "", fmt.Errorf("failed to read client bundle %s: %w", jsPath, err)
			}
		}

		return func() (string, error) {
			return string(js), nil
		}
	}

	renderer := render.NewCoreRenderer(jsf)

	mux := server.NewRouter(renderer)

	server.Start(mux, 8080)
}
