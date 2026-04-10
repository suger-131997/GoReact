package builder

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"os/exec"
)

//go:embed bundle.ts
var bundleTS string

type Builder struct {
	rootTmpl *template.Template
	envs     []string
}

func NewDevBuilder() *Builder {
	return &Builder{
		rootTmpl: template.Must(template.New("root").Parse(rootTemplate)),
		envs: []string{
			"NODE_ENV=development",
		},
	}
}

func NewProdBuilder() *Builder {
	return &Builder{
		rootTmpl: template.Must(template.New("root").Parse(rootTemplate)),
		envs: []string{
			"NODE_ENV=production",
		},
	}
}

func (b *Builder) Build(entryPoint string) (string, error) {
	cmd := exec.Command("node", "-e", bundleTS)
	cmd.Env = append(cmd.Env, b.envs...)

	var buf bytes.Buffer
	err := b.rootTmpl.Execute(&buf, map[string]interface{}{
		"EntryPoint": entryPoint,
	})
	if err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}
	cmd.Stdin = &buf

	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to run bundle.ts: %w, stderr: %s", err, stderr.String())
	}

	return out.String(), nil
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
