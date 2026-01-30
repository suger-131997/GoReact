package main

import (
	"goreact/internal/builder"
	"goreact/internal/render"
	"goreact/server"
)

func main() {
	b := builder.NewDevBuilder()
	cache := make(map[string]string)
	jsf := func(entryPoint string) func() (string, error) {
		return func() (string, error) {
			if js, ok := cache[entryPoint]; ok {
				return js, nil
			}

			js, err := b.Build(entryPoint)
			if err != nil {
				return "", err
			}
			cache[entryPoint] = js

			return js, nil
		}
	}

	renderer := render.NewCoreRenderer(jsf)

	mux := server.NewRouter(renderer)

	server.Start(mux, 3000)
}
