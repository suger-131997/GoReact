package main

import (
	"goreact/internal/builder"
	"goreact/internal/render"
	"goreact/server"
)

func main() {
	b := builder.NewDevBuilder()
	cache := make(map[string]builder.BuildResult)
	jsf := func(entryPoint string) func() (builder.BuildResult, error) {
		return func() (builder.BuildResult, error) {
			if br, ok := cache[entryPoint]; ok {
				return br, nil
			}

			br, err := b.Build(entryPoint)
			if err != nil {
				return br, err
			}
			cache[entryPoint] = br

			return br, nil
		}
	}

	renderer := render.NewCoreRenderer(jsf)

	mux := server.NewRouter(renderer)

	server.Start(mux, 3000)
}
