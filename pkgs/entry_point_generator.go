package pkgs

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"path"
	"path/filepath"
	"text/template"
)

type EntryPointGenerator struct {
	workdir  string
	rootTmpl *template.Template

	entryPoints map[string]struct{}
}

func NewEntryPointGenerator(workdir string) *EntryPointGenerator {
	return &EntryPointGenerator{
		workdir:     workdir,
		rootTmpl:    template.Must(template.New("root").Parse(entryPointTmpl)),
		entryPoints: make(map[string]struct{}),
	}
}

type EntryPointGeneratorContextKey struct{}

func WithEntryPointGenerator(ctx context.Context, generator *EntryPointGenerator) context.Context {
	return context.WithValue(ctx, EntryPointGeneratorContextKey{}, generator)
}

func EntryPointGeneratorFromContext(ctx context.Context) (*EntryPointGenerator, error) {
	value := ctx.Value(EntryPointGeneratorContextKey{})
	if value == nil {
		return nil, errors.New("entry point generator not found in context")
	}
	generator, ok := value.(*EntryPointGenerator)
	if !ok {
		return nil, errors.New("invalid entry point generator type in context")
	}
	return generator, nil
}

func (r *EntryPointGenerator) Generate(entryPoint string) error {
	if _, ok := r.entryPoints[entryPoint]; ok {
		return errors.New("entry point already exists")
	}

	r.entryPoints[entryPoint] = struct{}{}

	p := filepath.Join(r.workdir, entryPoint)

	err := os.MkdirAll(path.Dir(p), os.ModePerm)
	if err != nil {
		return err
	}

	f, err := os.Create(p)
	if err != nil {
		return err
	}
	defer f.Close()

	err = r.rootTmpl.Execute(f, map[string]interface{}{
		"EntryPoint": entryPoint,
	})
	if err != nil {
		return err
	}

	return nil
}

const entryPointTmpl = `
import { StrictMode } from "react";
import { createRoot } from 'react-dom/client'
import App from '~/{{ .EntryPoint }}'


createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <App {...(window.APP_PROPS || {})}/>
  </StrictMode>
)
`

func (r *EntryPointGenerator) GenerateEntryPointConfig() error {
	m := make(map[string]string, len(r.entryPoints))

	for entryPoint := range r.entryPoints {
		m[entryPoint] = filepath.Join(r.workdir, entryPoint)
	}

	b, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile("entries.gen.json", b, 0644)
	if err != nil {
		return err
	}

	return nil
}
