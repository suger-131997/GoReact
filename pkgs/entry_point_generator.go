package pkgs

import (
	"context"
	"errors"
	"os"
	"path"
	"path/filepath"
	"text/template"
)

type EntryPointGenerator struct {
	workdir  string
	rootTmpl *template.Template
}

func NewEntryPointGenerator(workdir string) *EntryPointGenerator {
	return &EntryPointGenerator{
		workdir:  workdir,
		rootTmpl: template.Must(template.New("root").Parse(entryPointTmpl)),
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
