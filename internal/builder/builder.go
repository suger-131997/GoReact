package builder

import (
	"bytes"
	"fmt"
	"html/template"

	esbuild "github.com/evanw/esbuild/pkg/api"
)

type Builder struct {
	baseOptions esbuild.BuildOptions

	rootTmpl *template.Template
}

func NewDevBuilder() *Builder {
	return &Builder{
		baseOptions: esbuild.BuildOptions{
			Bundle:            true,
			Write:             false,
			Tsconfig:          "./tsconfig.json",
			MinifyWhitespace:  false,
			MinifyIdentifiers: false,
			MinifySyntax:      false,
			Sourcemap:         esbuild.SourceMapInline,
			Define: map[string]string{
				"process.env.NODE_ENV": "\"development\"",
			},
		},
		rootTmpl: template.Must(template.New("root").Parse(rootTemplate)),
	}
}

func NewProdBuilder() *Builder {
	return &Builder{
		baseOptions: esbuild.BuildOptions{
			Bundle:            true,
			Write:             false,
			Tsconfig:          "./tsconfig.json",
			MinifyWhitespace:  true,
			MinifyIdentifiers: true,
			MinifySyntax:      true,
			Sourcemap:         esbuild.SourceMapNone,
			Define: map[string]string{
				"process.env.NODE_ENV": "\"production\"",
			},
		},
		rootTmpl: template.Must(template.New("root").Parse(rootTemplate)),
	}
}

func (b *Builder) copyBaseOptions() esbuild.BuildOptions {
	return b.baseOptions
}

func (b *Builder) Build(entryPoint string) (string, error) {
	var buf bytes.Buffer
	err := b.rootTmpl.Execute(&buf, map[string]interface{}{
		"EntryPoint": entryPoint,
	})
	if err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	opt := b.copyBaseOptions()
	opt.Stdin = &esbuild.StdinOptions{
		Contents:   buf.String(),
		ResolveDir: ".",
		Loader:     esbuild.LoaderTSX,
	}

	result := esbuild.Build(opt)

	if len(result.Errors) > 0 {
		fileLocation := "unknown"
		lineNum := "unknown"
		if result.Errors[0].Location != nil {
			fileLocation = result.Errors[0].Location.File
			lineNum = result.Errors[0].Location.LineText
		}
		return "", fmt.Errorf("%s <br>in %s <br>at %s", result.Errors[0].Text, fileLocation, lineNum)
	}

	return string(result.OutputFiles[0].Contents), nil
}

const rootTemplate = `
import { StrictMode } from "react";
import { createRoot } from 'react-dom/client'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import App from '~/{{ .EntryPoint }}'

const queryClient = new QueryClient()

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <QueryClientProvider client={queryClient}>
      <App {...(window.APP_PROPS || {})}/>
    </QueryClientProvider>
  </StrictMode>
)
`
